package pkg

import "github.com/redis/go-redis/v9"

type RedisConf struct {
	Addr     string
	Password string
	Db       int
}

func NewRedis(conf RedisConf) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.Db,
	})
	return rdb
}
