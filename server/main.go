package main

import (
	"fmt"
	"server/controller"
	mysql "server/dao"
	"server/entity"
	"server/middleWare"
	"server/utils"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	r := gin.Default()
	mysql.InitMysql()
	r.Use(middleWare.Cors())

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

	r.GET("/userList", func(c *gin.Context) {
		userList := controller.GetAllUser()
		c.JSON(200, gin.H{
			"data": userList,
		})
	})
	r.GET("/userInfo", func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		claims, err := utils.ParseToken(token)
		if err != nil {
			fmt.Println(err)
		}
		c.JSON(200, gin.H{
			"username": claims.Username,
			"password": claims.Password,
		})
	})

	r.Run(":8080")
}
