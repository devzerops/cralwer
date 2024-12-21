package crawler

import (
    "distributed-crawler/models"
    "distributed-crawler/storage"
    "github.com/gocolly/colly/v2"
    "log"
    "net/url"
    "strings"
)

type Crawler struct {
    Storage storage.Storage
}

func NewCrawler(storage storage.Storage) *Crawler {
    return &Crawler{Storage: storage}
}

func (c *Crawler) Crawl(startURL string) {
    // Create a new collector
    collector := colly.NewCollector()

    // Parse the start URL
    baseURL, err := url.Parse(startURL)
    if err != nil {
        log.Fatalf("Invalid start URL: %v", err)
    }

    // On every a element which has href attribute call callback
    collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        // Ignore URLs with fragments
        if strings.Contains(link, "#") {
            log.Printf("Ignoring link with fragment: %s", link)
            return
        }
        // Parse the link URL
        linkURL, err := url.Parse(link)
        if err != nil {
            log.Printf("Invalid link URL: %v", err)
            return
        }
        // Resolve the link URL to an absolute URL
        absoluteURL := baseURL.ResolveReference(linkURL).String()
        log.Printf("Found link: %s", absoluteURL)
        // Save the link to storage
        c.Storage.Save(models.URL{Address: absoluteURL, Priority: 1})
    })
    collector.OnHTML("body", func(e *colly.HTMLElement) {
        // parsing 
        content := e.Text
        log.Printf("Found content: %s", content)
        // Save the content to storage
        //c.Storage.SaveContent(models.Content{URL: e.Request.URL.String(), Data: content})
    })



    // Start scraping the URL
    collector.Visit(startURL)
}