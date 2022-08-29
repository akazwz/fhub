package middleware

import (
	"strings"
	"time"

	"github.com/akazwz/fhub/model/response"
	"github.com/akazwz/fhub/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.Request.Header.Get("Authorization")
		token := strings.Replace(bearerToken, "Bearer ", "", 1)
		if len(token) < 1 {
			response.Unauthorized(400, "未登录", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		// 解析 token
		claims, err := j.ParseToken(token)
		// 解析 token 错误
		if err != nil {
			if err == utils.TokenExpired {
				response.Unauthorized(400, "token已过期", c)
				c.Abort()
				return
			}
			response.Unauthorized(400, err.Error(), c)
			c.Abort()
			return
		}

		// 缓冲期内 可以 刷新 token(暂时不实现)
		if claims.ExpiresAt.Unix()-time.Now().Unix() < claims.BufferTime {

		}
		c.Set("uid", claims.UID)
		c.Next()
	}
}
