package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type config struct {
	Addr     string
	Port     int
	Password string
	Db       int
}

func Cache() *redis.Client {
	var conf config
	viper.UnmarshalKey("redis.cache", &conf)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
	})
	return client
}

func Db() *redis.Client {
	var conf config
	viper.UnmarshalKey("redis.db", &conf)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
		Password: conf.Password,
		DB:       conf.Db,
	})
	return client
}
