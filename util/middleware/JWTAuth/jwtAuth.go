package JWTAuth

import (
	jwt "ginhi/util/jwt"
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
)

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
		logging.Info("get token: ", token)

		j := jwt.NewJwt()
		// 解析token信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
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
		}
		//将解析的消息交给下一个路由处理
		c.Set("claims", claims)
		c.Next()
	}
}
