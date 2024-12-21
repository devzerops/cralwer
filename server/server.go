package server

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "distributed-crawler/storage"
)

// @POST("/crawl")
// @GET("/")
// @GET("/status")
// @GET("/next")
type Usage struct {
    Message      string `json:"message"`
    Endpoints    []struct {
        Method      string `json:"method"`
        Path        string `json:"path"`
        Description string `json:"description"`
    } `json:"endpoints"`
    ExampleUsage string `json:"example_usage"`
}

var redisStorage *storage.RedisStorage

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}


func init() {
    redisStorage = storage.NewRedisStorage()
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.Use(loggingMiddleware)
    router.HandleFunc("/crawl", CrawlHandler).Methods("POST")
    router.HandleFunc("/", GuideHandler).Methods("GET")
    router.HandleFunc("/status", StatusHandler).Methods("GET")
    router.HandleFunc("/next", ExtractHandler).Methods("GET")
    return router
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        //log.Printf("Started %s %s", r.Method, r.RequestURI)

        // Create a response writer wrapper to capture the status code
        rw := &responseWriter{w, http.StatusOK}
        next.ServeHTTP(rw, r)

        if rw.statusCode >= 200 && rw.statusCode < 300 {
            log.Printf("Completed %s %s with status %d in %v", r.Method, r.RequestURI, rw.statusCode, time.Since(start))
        } else {
            log.Printf("Failed %s %s with status %d in %v", r.Method, r.RequestURI, rw.statusCode, time.Since(start))
        }
    })
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}