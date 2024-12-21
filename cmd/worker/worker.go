package worker

import (
	"distributed-crawler/config"
	"distributed-crawler/crawler"
	"time"
    "strconv"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    serverApi := config.ServerAddress
	nextQuery := "next?count="
    count := 10
    delay := 10

    workerCrawler := crawler.NewWorkerCrawler(serverApi + nextQuery + strconv.Itoa(count))

    for {
        workerCrawler.Crawl()
        time.Sleep(time.Duration(delay) * time.Second)// Adjust the interval as needed
    }
}