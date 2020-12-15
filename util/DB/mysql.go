package db

import (
	"database/sql"
	"fmt"
	"ginhi/util/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"sync"
	"time"
)

var one sync.Once
var DB *gorm.DB

// NewConn is 创建数据库连接.
func InitDB() *gorm.DB {
	var err error
	one.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.GetString("mysql.user"),
			config.GetString("mysql.pwd"),
			config.GetString("mysql.host"),
			config.GetInt("mysql.port"),
			config.GetString("mysql.db"))
		// 连接池配置
		sqlDB, _ := sql.Open("mysql", dsn)
		if sqlDB != nil {
			sqlDB.SetMaxIdleConns(config.GetInt("mysql.maxIdle"))
			sqlDB.SetMaxOpenConns(config.GetInt("mysql.maxConns"))
			sqlDB.SetConnMaxLifetime(10 * time.Minute)
		}
		// 日志配置
		gormLog := logger.New(
			log.New(logrus.StandardLogger().Out, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: 1 * time.Second, // 慢 SQL
				LogLevel:      logger.Info,     // Loglevel 打印sql
				Colorful:      false,           // 禁用彩色打印
			},
		)
		// 初始化DB
		DB, err = gorm.Open(
			mysql.New(mysql.Config{
				Conn: sqlDB,
			}),
			&gorm.Config{
				Logger: gormLog,
				NamingStrategy: schema.NamingStrategy{
					SingularTable: true, // 使用单数表名
				},
				PrepareStmt: true, // 预编译
			})

	})
	if err != nil {
		logrus.Panic(err)
	}
	return DB
}
