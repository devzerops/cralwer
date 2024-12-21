package crawler

import (
    "distributed-crawler/models"
    "distributed-crawler/storage"
    "github.com/gocolly/colly/v2"
    "log"
)

type Crawler struct {
    Storage storage.Storage
}

func NewCrawler(storage storage.Storage) *Crawler {
    return &Crawler{Storage: storage}
}

func (c *Crawler) Crawl(url string) {
    // Create a new collector
    collector := colly.NewCollector()

    // On every a element which has href attribute call callback
    collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        log.Printf("Found link: %s", link)
        // Save the link to storage
        c.Storage.Save(models.URL{Address: link, Priority: 1})
    })

    // Start scraping the URL
    collector.Visit(url)
}