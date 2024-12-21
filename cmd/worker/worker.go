package worker

import (

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

