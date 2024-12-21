package config

import (
    "os"
    "log"
    "github.com/joho/godotenv"
)

const (
    MaxRequestsPerSite = 3
    RequestInterval    = 10 // seconds
)

var RedisAddress string

func init() {
    // .env 파일 로드
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    RedisAddress = getEnv("REDIS_ADDRESS", "localhost:6379")
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}