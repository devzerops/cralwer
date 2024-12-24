package storage

import (
    "github.com/go-redis/redis/v8"
    "context"
    "strings"
)

var ctx = context.Background()

// InitializeRedisClient initializes a Redis client
func InitializeRedisClient(addr string) *redis.Client {
    rdb := redis.NewClient(&redis.Options{
        Addr: addr,
    })
    return rdb
}

// AddMaliciousKeyword adds a malicious keyword to Redis if it doesn't already exist
func AddMaliciousKeyword(rdb *redis.Client, keyword string) error {
    exists, err := rdb.SIsMember(ctx, "malicious_keywords", keyword).Result()
    if (err != nil) {
        return err
    }
    if !exists {
        return rdb.SAdd(ctx, "malicious_keywords", keyword).Err()
    }
    return nil
}

// RemoveMaliciousKeyword removes a malicious keyword from Redis if it exists
func RemoveMaliciousKeyword(rdb *redis.Client, keyword string) error {
    exists, err := rdb.SIsMember(ctx, "malicious_keywords", keyword).Result()
    if err != nil {
        return err
    }
    if exists {
        return rdb.SRem(ctx, "malicious_keywords", keyword).Err()
    }
    return nil
}

// IsMaliciousURLByKeyword checks if the URL contains any malicious keywords stored in Redis
func IsMaliciousURLByKeyword(rdb *redis.Client, url string) (bool, error) {
    keywords, err := rdb.SMembers(ctx, "malicious_keywords").Result()
    if err != nil {
        return false, err
    }

    for _, keyword := range keywords {
        if strings.Contains(url, keyword) {
            return true, nil
        }
    }
    return false, nil
}