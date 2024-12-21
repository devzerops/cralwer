package worker

import (
	"encoding/json"
	"net/http"
	"distributed-crawler/internal/config"
	"distributed-crawler/internal/crawler"
	"distributed-crawler/internal/utils"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var workerCrawler *crawler.WorkerCrawler

func NewRouter() *mux.Router {
	serverApi := config.ServerAddress
	nextQuery := "next?count="
	count := 10
	delay := 10

	workerCrawler = crawler.NewWorkerCrawler(serverApi + nextQuery + strconv.Itoa(count))

	router := mux.NewRouter().StrictSlash(true)
	router.Use(utils.LoggingMiddleware)

	router.HandleFunc("/health", HealthHandler).Methods("GET")
	router.HandleFunc("/status", StatusHandler).Methods("GET")
	router.HandleFunc("/fetch-url", FetchURLHandler).Methods("GET")

	go func() {
		for {
			workerCrawler.Crawl()
			time.Sleep(time.Duration(delay) * time.Second) // Adjust the interval as needed
		}
	}()

	return router
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := map[string]string{
		"status": "running",
	}
	json.NewEncoder(w).Encode(status)
}

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
