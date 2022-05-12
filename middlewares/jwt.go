package middlewares

import (
	"SilicomAPPv0.3/common"
	"SilicomAPPv0.3/response"
	"github.com/gin-gonic/gin"
)

// JwtAuth JWT认证中间件
func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.Failed("未登录或非法访问", c)
			c.Abort()
			return
		}
		username, err := common.VerifyToken(token)
		if err != nil {
			response.Failed("登录已过期，请重新登录", c)
			c.Abort()
			return
		}
		c.Set("username", username) // 在Context中添加用户名信息
		c.Next()
	}
}
