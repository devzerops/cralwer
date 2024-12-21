package server

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "time"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.Use(loggingMiddleware)
    router.HandleFunc("/crawl", CrawlHandler).Methods("POST")
    router.HandleFunc("/", UsageHandler).Methods("GET")
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
    usage := `
    Welcome to the Distributed Crawler API!

    Available endpoints:
    POST /crawl - Start crawling a URL. Example request body: {"url": "http://example.com"}

    Example usage:
    curl -X POST -H "Content-Type: application/json" -d '{"url":"http://example.com"}' http://localhost:8080/crawl
    `
    fmt.Fprintln(w, usage)
}