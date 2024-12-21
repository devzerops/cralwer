package storage

import (
	"context"
	"distributed-crawler/internal/config"
	"distributed-crawler/internal/models"
	"encoding/json"
	"log"

	"github.com/go-redis/redis/v8"
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

func (s *RedisStorage) QueueLength() int64 {
	length, err := s.client.DBSize(s.ctx).Result()
	if err != nil {
		log.Printf("Error getting queue length from Redis: %v", err)
		return 0
	}
	return length
}

func (s *RedisStorage) NextURL() (models.URL, error) {
	keys, err := s.client.Keys(s.ctx, "*").Result()
	if err != nil {
		return models.URL{}, err
	}
	if len(keys) == 0 {
		return models.URL{}, nil
	}
	key := keys[0]
	data, err := s.client.Get(s.ctx, key).Result()
	if err != nil {
		return models.URL{}, err
	}
	var url models.URL
	err = json.Unmarshal([]byte(data), &url)
	if err != nil {
		return models.URL{}, err
	}
	err = s.client.Del(s.ctx, key).Err()
	if err != nil {
		return models.URL{}, err
	}
	return url, nil
}
