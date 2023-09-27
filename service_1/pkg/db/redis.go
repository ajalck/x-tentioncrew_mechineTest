package db

import (
	"github.com/ajalck/service_1/pkg/config"
	"github.com/go-redis/redis/v8"
)

func ConnectRedis(config *config.Config) *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr: config.RedisURL,
		})
	return rdb
}