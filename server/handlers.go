package server

import (
	"distributed-crawler/crawler"
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

	c := crawler.NewCrawler(redisStorage)
	c.Crawl(req.URL)

	w.WriteHeader(http.StatusOK)
}
