package crawler

import (
    //"distributed-crawler/models"
    "distributed-crawler/storage"
)

type Crawler struct {
    Storage storage.Storage
}

func NewCrawler(storage storage.Storage) *Crawler {
    return &Crawler{Storage: storage}
}

func (c *Crawler) Crawl(url string) {
    // Implement crawling logic here
}
