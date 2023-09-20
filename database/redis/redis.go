package redis

import (
	"fmt"
	"foods-spider-server/pkg"
	"github.com/go-redis/redis"
)

func InitRedisDB() (rdb *redis.Client, err error) {
	if pkg.Config.RedisConfig == nil {
		panic("[RedisConfig] Unavailable configuration")
	}
	DSN := fmt.Sprintf("%s:%v", pkg.Config.RedisConfig.Host, pkg.Config.RedisConfig.Port)
	rdb = redis.NewClient(&redis.Options{
		Addr:     DSN,
		Password: pkg.Config.RedisConfig.Password, // 没有密码设置 no password set
		DB:       pkg.Config.RedisConfig.Db,       // 默认 db use default DB
	})
	sc := rdb.Ping()
	if sc.Err() != nil {
		return nil, sc.Err()
	}
	return //tips:这里的返回值中有有变量名，所以直接return就行了！
}
