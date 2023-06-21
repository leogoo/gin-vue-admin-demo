package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"server/controller"
	mysql "server/dao"
	"server/entity"
	"server/middleWare"
	"server/routers"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	mysql.InitMysql()
	r.Use(middleWare.Cors())

	r.GET("/test", func(ctx *gin.Context) {
		fmt.Println(1111)
		pusher := ctx.Writer.Pusher()
		fmt.Println("pusher", pusher)
		if pusher != nil {
			fmt.Println(2222)
			if err := pusher.Push("gin.log", nil); err != nil {
				log.Printf("push error %v", err)
			}
		}
		ctx.JSON(200, gin.H{
			"status": "success",
		})
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

	r.POST("/user/add", func(ctx *gin.Context) {
		json := entity.User{}
		ctx.BindJSON(&json)

		controller.CreateUser(json)
	})

	r.Use(middleWare.JWT())
	routers.SetUserRouter(r)
	r.Run(":8080")
}
