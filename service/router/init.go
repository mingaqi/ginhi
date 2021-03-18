package router

import (
	_ "ginhi/docs"
	"ginhi/service/router/handle"
	"ginhi/service/router/reverse"
	middle "ginhi/util/middleware/Cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option func(engine *gin.Engine)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	gin.DefaultWriter = logrus.StandardLogger().Out
	gin.DefaultErrorWriter = logrus.StandardLogger().Out
	gin.ForceConsoleColor()
	Router := gin.New()
	Router.Use(middle.Cors(), gin.Logger(), gin.Recovery())
	setHandler(Router)
	return Router
}

func setHandler(router *gin.Engine) {

	/*------------------------------swagger----------------------------------------*/
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/*------------------------------业务路由----------------------------------------*/

	// 分组
	vv := router.Group("/v1")
	{
		vv.GET("/user/:name", handle.UserSave)
		vv.GET("/user", handle.UserSaveByQuery)
		vv.POST("/user/register", handle.UserRegister)
	}
	// 重定向
	router.GET("/redirect", handle.Redirect)
	router.POST("/jsontest", handle.JsonTest)

	// gin反向代理
	router.GET("/reverseProxy", reverse.ReverseProxy)

}
