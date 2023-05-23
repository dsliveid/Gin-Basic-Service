package util

import (
	"Gin-Basic-Service/global"
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitDB() {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}
	// 设置全局连接
	global.DB = db

}

func GetDB() (db *gorm.DB, err error) {
	var dialector gorm.Dialector
	switch global.Conf.Database.Driver {
	case global.DriverNameMySQL:
		dialector = mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%s&loc=%s",
			global.Conf.Database.Username, global.Conf.Database.Password, global.Conf.Database.Host, global.Conf.Database.Database, "utf8mb4", "True", "Local"))
	default:

	}
	// 如果没有获取到dialector则退出
	if dialector == nil {
		return
	}
	// 打开 GORM 数据库连接
	// PrepareStmt = false 创建并缓存预编译语句，需要关闭才能检测数据库连接地址有效性
	if db, err = gorm.Open(dialector, &gorm.Config{PrepareStmt: false, Logger: getLogger()}); err != nil {
		return
	}

	// 数据库连接池配置
	var sqlDB *sql.DB
	if sqlDB, err = db.DB(); err != nil {
		return
	}
	sqlDB.SetMaxIdleConns(20)           // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(1e3)          // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // 设置连接可复用的最大时间

	return
}

func getLogger() logger.Interface {
	gormLogger := logger.New(
		log.New(os.Stdout, "", log.LstdFlags),
		logger.Config{
			SlowThreshold: 3 * time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Warn,     // 日志记录级别
			Colorful:      true,            // 是否启用彩色打印
		},
	)

	return gormLogger
}
