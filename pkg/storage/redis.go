package storage

import (
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(addr, password string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() error {
	return redisClient.Close()
}
