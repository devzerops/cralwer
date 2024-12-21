package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "time"
    "distributed-crawler/storage"
)

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

func init() {
    redisStorage = storage.NewRedisStorage()
}

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.Use(loggingMiddleware)
    router.HandleFunc("/crawl", CrawlHandler).Methods("POST")
    router.HandleFunc("/", UsageHandler).Methods("GET")
    router.HandleFunc("/status", StatusHandler).Methods("GET")
    return router
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        log.Printf("Started %s %s", r.Method, r.RequestURI)
        next.ServeHTTP(w, r)
        log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
    })
}

func UsageHandler(w http.ResponseWriter, r *http.Request) {
    data, err := ioutil.ReadFile("config/usage.json")
    if err != nil {
        http.Error(w, "Unable to read usage file", http.StatusInternalServerError)
        return
    }

    var usage Usage
    err = json.Unmarshal(data, &usage)
    if err != nil {
        http.Error(w, "Unable to parse usage file", http.StatusInternalServerError)
        return
    }

    fmt.Fprintf(w, "%s\n\nAvailable endpoints:\n", usage.Message)
    for _, endpoint := range usage.Endpoints {
        fmt.Fprintf(w, "%s %s - %s\n", endpoint.Method, endpoint.Path, endpoint.Description)
    }
    fmt.Fprintf(w, "\nExample usage:\n%s\n", usage.ExampleUsage)
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
    queueLength := redisStorage.QueueLength()
    fmt.Fprintf(w, "Total items in queue: %d\n", queueLength)
}