package main

import (
	"foods-spider-server/database/mysql"
	"foods-spider-server/internal/engine"
	"foods-spider-server/pkg"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"log"
)

var err error

// RedisDB redis 实例对象
var RedisDB *redis.Client

// MysqlDB mysql 实例对象
var MysqlDB *gorm.DB

// Logger 日志实例对象！
var Logger *zap.Logger

func runApp() {
	pkg.InitConfig()
	/*
		初始化日志
	*/
	Logger, err = pkg.InitLogger()
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Logger Ready ! %v \n", Logger)
	// 替换全局的	zap.L() => logger
	zap.ReplaceGlobals(Logger)
	/*
		连接 redis
	*/
	//RedisDB, err = redisIns.InitRedisDB()
	//if err != nil {
	//	log.Fatal(err.Error())
	//}
	log.Printf("RedisDB Ready ! %v \n", RedisDB)
	/*
		连接 mysql
	*/
	MysqlDB, err = mysql.InitSqlxDB()
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	log.Printf("MysqlDB Ready ! %v \n", MysqlDB)
	/*
		初始化翻译器！
	*/
	err = pkg.InitLocaleTrans("zh")
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	log.Println("LocaleTrans Ready ! ")
	/*
		初始化 雪花id 引擎~
	*/
	err = pkg.InitSnowflake(pkg.Config.App.OnlineTime, pkg.Config.MachineID)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	log.Println("Snowflake Ready !")
	engine.CreateServer()
}

func main() {
	runApp()
}
