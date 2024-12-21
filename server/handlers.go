package server

import (
	"distributed-crawler/crawler"
	"distributed-crawler/storage"
	"encoding/json"
	"net/http"
)

type CrawlRequest struct {
	URL string `json:"url"`
}

func CrawlHandler(w http.ResponseWriter, r *http.Request) {
	var req CrawlRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	storage := storage.NewRedisStorage()
	c := crawler.NewCrawler(storage)
	c.Crawl(req.URL)

	w.WriteHeader(http.StatusOK)
}
