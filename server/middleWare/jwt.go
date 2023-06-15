package middleWare

import (
	"net/http"
	"time"

	"server/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		var code int = 0
		var msg string = ""
		if token == "" {
			code = -3
			msg = "未登录"
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = -1
				msg = "token解析失败"
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = -2
				msg = "token过期"
			}
		}
		if code != 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  msg,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
