package db

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// Mysql is 数据库配置.
type Mysql struct {
	Host         string `json:"host" toml:"host" description:"数据库地址"`
	Port         int    `json:"port" toml:"port" description:"数据库端口"`
	Username     string `json:"username" toml:"username" description:"数据库访问用户名"`
	Password     string `json:"password" toml:"password" description:"数据库访问密码"`
	Database     string `json:"database" toml:"database" description:"数据库名称"`
	MaxOpenConns int    `json:"maxOpenConns" toml:"maxOpenConns" description:"数据库最大连接数"`
	MaxIdleConns int    `json:"maxIdleConns" toml:"maxIdleConns" description:"数据库最大空闲连接数"`

	db *gorm.DB
}

// NewConn is 创建数据库连接.
func (p *Mysql) NewConn() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", p.Username, p.Password, p.Host, p.Port, p.Database)
	// 连接池配置
	sqlDB, err := sql.Open("mysql", dsn)
	if sqlDB != nil {
		sqlDB.SetMaxIdleConns(p.MaxIdleConns)
		sqlDB.SetMaxOpenConns(p.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(3 * time.Minute)
	}

	// 日志配置
	gormLog := logger.New(
		log.New(logrus.StandardLogger().Out, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: 500 * time.Millisecond, // 慢 SQL
			LogLevel:      logger.Info,            // Loglevel 打印sql
			Colorful:      false,                  // 禁用彩色打印
		},
	)
	// 初始化DB
	db, err := gorm.Open(
		mysql.New(mysql.Config{
			Conn: sqlDB, // 使用原生sql连接池
		}),
		&gorm.Config{
			Logger: gormLog,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
		})
	if err != nil {
		logrus.Fatal("数据库初始化错误:", err)
		return nil, err
	}
	return db, nil
}
