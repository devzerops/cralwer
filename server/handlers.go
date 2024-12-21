package server

import (
	"distributed-crawler/crawler"
	"encoding/json"
	"net/http"
	"io/ioutil"
	"fmt"
	"strconv"
	"distributed-crawler/models"
)

type CrawlRequest struct {
	URL string `json:"url"`
}

// @POST /crawl
// CrawlHandler handles the "/crawl" endpoint, starting the crawling process for the given URL.
func CrawlHandler(w http.ResponseWriter, r *http.Request) {
    var req CrawlRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    c := crawler.NewCrawler(redisStorage)
    c.Crawl(req.URL)

    w.WriteHeader(http.StatusOK)
}

// @GET /
// UsageHandler handles the "/" endpoint, providing API usage information.
func GuideHandler(w http.ResponseWriter, r *http.Request) {
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

// @GET /status
// StatusHandler handles the "/status" endpoint, providing the total number of items in the Redis queue.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
    queueLength := redisStorage.QueueLength()
    fmt.Fprintf(w, "Total items in queue: %d\n", queueLength)
}

// @GET /next
// ExtractHandler handles the "/next" endpoint, extracting the next URL(s) from the Redis queue.
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
        url, err := redisStorage.NextURL()
        if err != nil {
            http.Error(w, "Unable to next URL from queue", http.StatusInternalServerError)
            return
        }
        if url.Address == "" {
            break
        }
        urls = append(urls, url)
    }

    json.NewEncoder(w).Encode(urls)
}