package utils

import (
    "log"
    "net/http"
    "time"
)


func IsFresh(url string) bool {
    // Implement freshness check logic here
    return true
}

type ResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func (rw *ResponseWriter) WriteHeader(code int) {
    rw.StatusCode = code
    rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()

        // Create a response writer wrapper to capture the status code
        rw := &ResponseWriter{w, http.StatusOK}
        next.ServeHTTP(rw, r)

        if rw.StatusCode >= 200 && rw.StatusCode < 300 {
            log.Printf("Completed %s %s with status %d in %v", r.Method, r.RequestURI, rw.StatusCode, time.Since(start))
        } else {
            log.Printf("Failed %s %s with status %d in %v", r.Method, r.RequestURI, rw.StatusCode, time.Since(start))
        }
    })
}