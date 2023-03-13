package middleweare

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/poeticalcode/gim/internal/tool"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")
		user := c.Query("userId")
		userId, err := strconv.Atoi(user)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "您userId不合法",
			})
			c.Abort()
			return
		}
		if token == "" {
			c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "请登录",
			})
			c.Abort()
			return
		} else {
			claims, err := tool.ParseToken(token)
			if err != nil {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "token失效",
				})
				c.Abort()
				return
			} else if time.Now().Unix() > claims.ExpiresAt {
				err = tool.TokenExpired
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "授权已过期",
				})
				c.Abort()
				return
			}

			if claims.UserID != uint(userId) {
				c.JSON(http.StatusUnauthorized, map[string]string{
					"message": "您的登录不合法",
				})
				c.Abort()
				return
			}

			fmt.Println("token认证成功")
			c.Next()
		}
	}
}
