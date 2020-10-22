package main

import (
	"ginhi/service/initRouter"
	"ginhi/util/config"
	"ginhi/util/logs"
	"github.com/sirupsen/logrus"
)

func main() {

	// 日志
	logs.InitLogs("./logs/ginhi.log")

	// 路由
	router := initRouter.InitRouter()

	err := router.Run(":" + config.GetViper().GetString("server.port"))

	logrus.Fatal(err)

}
