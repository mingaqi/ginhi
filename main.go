package main

import (
	"ginhi/service/router"
	"ginhi/util/config"
	"ginhi/util/logs"
	"github.com/gin-contrib/pprof"
	"github.com/sirupsen/logrus"
)

func main() {

	// 日志
	logs.InitLogs("./logs/ginhi%s.log")

	// 路由
	router := router.InitRouter()
	pprof.Register(router, "/de/pprof")
	err := router.Run(":" + config.GetViper().GetString("server.port"))

	logrus.Fatal(err)

}
