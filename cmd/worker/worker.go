package worker

import (
	"distributed-crawler/crawler"
	"time"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    serverApi := "http://localhost:8080/"
	
	query := "next?count="

    workerCrawler := crawler.NewWorkerCrawler(serverApi+query)

    for {
        workerCrawler.Crawl()
        time.Sleep(10 * time.Second) // Adjust the interval as needed
    }
}