package config

import (
	"flag"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"sync"
)

var (
	v    *viper.Viper
	once sync.Once
)

func GetViper() *viper.Viper {
	once.Do(func() {
		var filePath string
		flag.StringVar(&filePath, "config", "config/config.yaml", "指定配置文件的路径")
		// 解析配置文件
		flag.Parse()

		v = viper.GetViper()
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.SetConfigFile(filePath)
		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logrus.Fatal("配置文件未找到")
				os.Exit(1)
			}
		}
		switch v.GetString("mode") {
		case "dev":
			v.SetConfigFile("config/config-dev.yaml")
		case "test":
			v.SetConfigFile("config/config-test.yaml")
		case "prod":
			v.SetConfigFile("config/config-prod.yaml")
		default:
			logrus.Fatal("配置文件未找到")

		}
		err := v.ReadInConfig()
		if err != nil {
			logrus.Fatal("配置文件读取出错")
		}
	})
	return v
}

func GetInt(tag string) int {
	return v.GetInt(tag)
}
func GetString(tag string) string {
	return v.GetString(tag)
}
