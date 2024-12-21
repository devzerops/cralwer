package config

import (
    "log"
    "os"
    "gopkg.in/yaml.v3"
    "strconv"
)

const (
    MaxRequestsPerSite = 3
    RequestInterval    = 10 // seconds
)

var RedisAddress string
var ServerAddress string
var WorkerAddress string

type Config struct {
    Redis struct {
        Address string `yaml:"address"`
        Port    int    `yaml:"port"`
    } `yaml:"redis"`
    Server struct {
        Address  string `yaml:"address"`
        Port     int    `yaml:"port"`
        Protocol string `yaml:"protocol"`
        Env      string `yaml:"env"`
    } `yaml:"server"`
    Worker struct {
        Address  string `yaml:"address"`
        Port     int    `yaml:"port"`
        Protocol string `yaml:"protocol"`
        Env      string `yaml:"env"`
        Monitoring struct {
            Timeout         int `yaml:"timeout"`
            RequestInterval int `yaml:"RequestInterval"`
        } `yaml:"monitoring"`
        Crawling struct {
            RequestInterval int `yaml:"RequestInterval"`
            RequestPerBatch int `yaml:"RequestPerBatch"`
        } `yaml:"crawling"`
    } `yaml:"worker"`
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
    WorkerAddress = getEnv("WORKER_ADDRESS", config.Worker.Protocol + "://" + config.Worker.Address + ":" + strconv.Itoa(config.Worker.Port))
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}