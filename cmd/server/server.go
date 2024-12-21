package server

import (
	"distributed-crawler/internal/storage"
	"distributed-crawler/internal/utils"
	"github.com/gorilla/mux"
)

// @POST("/add-url")
// @GET("/")
// @GET("/status")
// @GET("/next")
// @GET("/health")
type Usage struct {
    Message   string `json:"message"`
    Endpoints []struct {
        Method      string `json:"method"`
        Path        string `json:"path"`
        Description string `json:"description"`
    } `json:"endpoints"`
    ExampleUsage string `json:"example_usage"`
}

var redisStorage *storage.RedisStorage

func init() {
    redisStorage = storage.NewRedisStorage()
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.Use(utils.LoggingMiddleware)
    router.HandleFunc("/add-url", CrawlHandler).Methods("POST")
    router.HandleFunc("/", GuideHandler).Methods("GET")
    router.HandleFunc("/status", StatusHandler).Methods("GET")
    router.HandleFunc("/next-url", NextURLHandler).Methods("GET")
    router.HandleFunc("/health", HealthHandler).Methods("GET")
    return router
}

