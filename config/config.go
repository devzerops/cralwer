package config

import (
    "log"
    "os"
    "gopkg.in/yaml.v2"
    "strconv"
)

var RedisAddress string
var ServerAddress string

type Config struct {
    Redis struct {
        Address string `yaml:"address"`
        Port    int    `yaml:"port"`
    } `yaml:"redis"`
    Server struct {
        Address  string `yaml:"address"`
        Port     int    `yaml:"port"`
        Protocol string `yaml:"protocol"`
    } `yaml:"server"`
}

func init() {
    // config.yaml 파일 로드
    configFile, err := os.ReadFile("config.yaml")
    if err != nil {
        log.Fatalf("Error reading config.yaml file: %v", err)
    }

    var config Config
    err = yaml.Unmarshal(configFile, &config)
    if err != nil {
        log.Fatalf("Error parsing config.yaml file: %v", err)
    }
    
    RedisAddress = getEnv("REDIS_ADDRESS", config.Redis.Address + ":" + strconv.Itoa(config.Redis.Port))
    ServerAddress = getEnv("SERVER_ADDRESS", config.Server.Protocol + "://" + config.Server.Address + ":" + strconv.Itoa(config.Server.Port))
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}