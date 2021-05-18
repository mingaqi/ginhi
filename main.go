package main

import (
	"context"
	"ginhi/service/router"
	"ginhi/util/config"
	"ginhi/util/logs"
	"github.com/gin-contrib/pprof"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// 日志
	logs.InitLogs("./logs/ginhi%s.log")

	// 路由
	router := router.InitRouter()
	pprof.Register(router, "/de/pprof")

	// ------  启动以及退出---------
	server := &http.Server{
		Addr:    ":" + config.GetViper().GetString("server.port"),
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	//监听信号量退出程序 kill -2 is syscall.SIGINT
	signal.Notify(signalChan, syscall.SIGINT)
	s := <-signalChan
	logrus.Info("退出信号:" + s.String())

	// 调用shutdown退成server
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorln(err)
	}
	logrus.Info("server 退出")

}
