package worker

import (
	"encoding/json"
	"net/http"
)

//@GET /health
// HealthHandler is a simple health check handler
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

//@GET /status
// StatusHandler is a simple status check handler
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "running",
	}
	json.NewEncoder(w).Encode(status)
}

//@GET /fetch-url
// FetchURLHandler fetches URLs from the server and crawls them
func FetchURLHandler(w http.ResponseWriter, r *http.Request) {
	urls, err := workerCrawler.FetchURLs()
	if err != nil {
		http.Error(w, "Failed to fetch URLs", http.StatusInternalServerError)
		return
	}

	for _, urlObj := range urls {
		workerCrawler.CrawlURL(urlObj.Address)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URLs fetched and crawled"))
}
