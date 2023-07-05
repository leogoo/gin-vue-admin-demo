package routers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"server/controller"
	dao "server/dao"
	"server/entity"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetUserRouter(r *gin.Engine) {

	r.POST("/api/user/add", func(ctx *gin.Context) {
		json := entity.User{}
		ctx.BindJSON(&json)

		err := controller.CreateUser(json)
		if err != nil {
			fmt.Fprint(gin.DefaultWriter, err)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": gin.H{},
		})
	})
	r.GET("/api/userList", func(c *gin.Context) {
		exists, err := dao.MyRedis.Exists(context.Background(), "userList").Result()
		// 判断 key 是否存在
		if err != nil {
			utils.Log(c)(err.Error())
		} else {
			if exists == 1 {
				usersListJSONFromRedis, err := dao.MyRedis.Get(context.Background(), "userList").Bytes()
				if err != nil {
					utils.Log(c)("get userlist failed", zap.String("error", err.Error()))
				}
				var userList []*entity.User
				err = json.Unmarshal(usersListJSONFromRedis, &userList)
				if err != nil {
					utils.Log(c)("get userlist failed", zap.String("error", err.Error()))
					panic(err)
				}
				c.JSON(200, gin.H{
					"data": userList,
				})
				return
			}
		}
		userList := controller.GetAllUser()
		userListJson, err := json.Marshal(userList)
		if err != nil {
			utils.Log(c)("json.Marshal failed")
			panic(err)
		}
		err = dao.MyRedis.Set(context.Background(), "userList", userListJson, 100*time.Second).Err()
		if err != nil {
			utils.Log(c)("set user list failed", zap.String("error", err.Error()))
		}
		c.JSON(200, gin.H{
			"data": userList,
		})
	})
	r.GET("/api/userInfo", func(c *gin.Context) {
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
