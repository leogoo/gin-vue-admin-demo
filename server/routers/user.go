package routers

import (
	"fmt"
	"server/controller"
	"server/utils"

	"github.com/gin-gonic/gin"
)

func SetUserRouter(r *gin.Engine) {

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
}
