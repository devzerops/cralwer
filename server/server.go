package server

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "strconv"
    "github.com/gorilla/mux"
    "time"
    "distributed-crawler/storage"
    "distributed-crawler/models"
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
    router.HandleFunc("/", UsageHandler).Methods("GET")
    router.HandleFunc("/status", StatusHandler).Methods("GET")
    router.HandleFunc("/extract", ExtractHandler).Methods("GET")
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

func ExtractHandler(w http.ResponseWriter, r *http.Request) {
    countStr := r.URL.Query().Get("count")
    count, err := strconv.Atoi(countStr)
    if err != nil || count <= 0 {
        http.Error(w, "Invalid count", http.StatusBadRequest)
        return
    }

    queueLength := redisStorage.QueueLength()
    if int64(count) > queueLength {
        http.Error(w, fmt.Sprintf("Requested count %d exceeds total items in queue %d", count, queueLength), http.StatusBadRequest)
        return
    }

    var urls []models.URL
    for i := 0; i < count; i++ {
        url, err := redisStorage.ExtractURL()
        if err != nil {
            http.Error(w, "Unable to extract URL from queue", http.StatusInternalServerError)
            return
        }
        if url.Address == "" {
            break
        }
        urls = append(urls, url)
    }

    json.NewEncoder(w).Encode(urls)
}