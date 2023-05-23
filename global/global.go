package global

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var Conf *Config
var DB *gorm.DB
var Router *gin.Engine

// RouterPrefix 路由前缀
const RouterPrefix = "/api/custom/routing"

// 支持的数据库驱动名称
const (
	DriverNameSQLite    = "sqlite"    // SQLite 数据库（github.com/glebarez/sqlite）
	DriverNameSQLite3   = "sqlite3"   // SQLite 数据库（gorm.io/driver/sqlite）
	DriverNameMySQL     = "mysql"     // MySQL 数据库
	DriverNamePostgres  = "postgres"  // PostgreSQL 数据库
	DriverNameSqlServer = "sqlserver" // SQL Server 数据库
	DriverNameOracle    = "oracle"    // Oracle 数据库
)

// 返回码常量
const (
	SuccessCode    = 0 // 成功
	ErrorCode      = 1 // 错误
	NotFoundCode   = 2 // 未找到
	Unauthorized   = 3 // 未授权
	InvalidRequest = 4 // 无效请求
)
