package mysql

import (
	"fmt"
	"foods-spider-server/pkg"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	logger.Config{
		SlowThreshold:             time.Second,   // 慢 SQL 阈值
		LogLevel:                  logger.Silent, // 日志级别
		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  false,         // 禁用彩色打印
	},
)

func InitSqlxDB() (*gorm.DB, error) {
	if pkg.Config.MysqlConfig == nil {
		panic("[MysqlConfig] Unavailable configuration")
	}
	dbUser := pkg.Config.MysqlConfig.Username
	dbPassword := pkg.Config.MysqlConfig.Password
	dbName := pkg.Config.MysqlConfig.DbName
	dbURL := pkg.Config.MysqlConfig.Host
	dbPort := pkg.Config.MysqlConfig.Port
	loc := pkg.Config.MysqlConfig.LocalTime
	sqlDSN := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=%s", dbUser, dbPassword, dbURL, dbPort, dbName, url.QueryEscape(loc))
	db, err := gorm.Open(mysql.Open(sqlDSN), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}
	originDB, _ := db.DB()
	originDB.SetMaxOpenConns(pkg.Config.MysqlConfig.MaxOpenConnes)
	originDB.SetMaxIdleConns(pkg.Config.MysqlConfig.MaxIdleConnes)
	return db, nil
}
