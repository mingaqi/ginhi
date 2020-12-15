package cjwt

import (
	"github.com/gin-gonic/gin"
)

//  token认证
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token") // 将token放入header中进行接收
		if token == "" {
			c.JSON(200, gin.H{
				"code":   400,
				"status": -1,
				"msg":    "请求未携带token，无权限访问",
			})
			c.Abort()
			return
		}

		// 解析token信息
		claims, err := ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				c.JSON(200, gin.H{
					"code":   400,
					"status": -1,
					"msg":    "授权过期",
				})
				c.Abort()
				return
			}
			c.JSON(200, gin.H{
				"code":   400,
				"status": -1,
				"msg":    err.Error(),
			})
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
