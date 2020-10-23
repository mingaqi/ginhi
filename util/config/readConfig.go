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
		flag.StringVar(&filePath, "config", "config/config.toml", "指定配置文件的路径")
		// 解析配置文件
		flag.Parse()

		v = viper.GetViper()
		v.SetConfigName("config")
		v.SetConfigType("toml")
		v.SetConfigFile(filePath)
		if err := v.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				logrus.Fatal("配置文件未找到")
				os.Exit(1)
			}
		}
		if v.GetString("mode") == "dev" {
			v.SetConfigFile("config/config-dev.toml")
			v.ReadInConfig()
		}
		if v.GetString("mode") == "test" {
			v.SetConfigFile("config/config-test.toml")
			v.ReadInConfig()
		}
		if v.GetString("mode") == "prod" {
			v.SetConfigFile("config/config-prod.toml")
			v.ReadInConfig()
		}
	})
	/*var c server
	err := v.UnmarshalKey("server", &c)*/
	return v
}

type server struct {
	Port int
}
