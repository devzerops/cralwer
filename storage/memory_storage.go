package storage

import (
    "context"
    "distributed-crawler/config"
    "distributed-crawler/models"
    "github.com/go-redis/redis/v8"
    "encoding/json"
    "log"
)

type RedisStorage struct {
    client *redis.Client
    ctx    context.Context
}

func NewRedisStorage() *RedisStorage {
    ctx := context.Background()
    client := redis.NewClient(&redis.Options{
        Addr: config.RedisAddress,
    })

    return &RedisStorage{
        client: client,
        ctx:    ctx,
    }
}

func (s *RedisStorage) Save(url models.URL) {
    data, err := json.Marshal(url)
    if err != nil {
        log.Printf("Error marshalling URL: %v", err)
        return
    }
    err = s.client.Set(s.ctx, url.Address, data, 0).Err()
    if err != nil {
        log.Printf("Error saving URL to Redis: %v", err)
    }
}

func (s *RedisStorage) Exists(url string) bool {
    exists, err := s.client.Exists(s.ctx, url).Result()
    if err != nil {
        log.Printf("Error checking URL existence in Redis: %v", err)
        return false
    }
    return exists > 0
}