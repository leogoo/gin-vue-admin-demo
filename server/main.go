package main

import (
	"fmt"
	"os"
	rabbitmq "server/amqp"
	"server/controller"
	dao "server/dao"
	"server/middleWare"
	"server/routers"
	"server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// 创建文件输出器
	file, _ := os.Create("logs/app.log")
	defer file.Close()

	// 创建 zapcore.WriteSyncer 对象
	writeSyncer := zapcore.AddSync(file)

	// 创建 zapcore.Encoder 对象
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

	// 创建 zapcore.Core 对象
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 创建 zap logger 对象
	logger := zap.New(core)

	r := gin.Default()
	dao.InitMysql()
	dao.InitRedis()
	rabbitmq.InitAmqp()

	r.Use(middleWare.Cors())
	// 注册 zap 中间件
	r.Use(func(c *gin.Context) {
		// 将 zap logger 对象注入 gin 的上下文中
		c.Set("logger", logger)

		// 继续处理请求
		c.Next()
	})

	// 登录，校验用户名和密码，生成token
	r.POST("/auth", func(ctx *gin.Context) {
		json := User{}
		ctx.BindJSON(&json)
		isValidUser := controller.CheckAuth(json.Username, json.Password)
		token, err := utils.GenerateToken(json.Username, json.Password)
		if err != nil {
			fmt.Println(err)
		}
		ctx.JSON(200, gin.H{
			"body":  json,
			"data":  isValidUser,
			"token": token,
		})
	})

	// 静态资源托管
	r.StaticFS("/static", gin.Dir("./upload", true))

	// r.Use(middleWare.JWT())
	routers.SetUserRouter(r)
	routers.SetMediaRouter(r)

	r.Run(":8080")
}
