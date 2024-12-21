package crawler

import (
    "distributed-crawler/models"
    "encoding/json"
    "log"
    "net/http"
    "net/url"
    "strings"
    "github.com/gocolly/colly/v2"
)

type WorkerCrawler struct {
    APIURL string
}

func NewWorkerCrawler(apiURL string) *WorkerCrawler {
    return &WorkerCrawler{APIURL: apiURL}
}

func (wc *WorkerCrawler) FetchURLs() ([]models.URL, error) {
    resp, err := http.Get(wc.APIURL)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var urls []models.URL
    if err := json.NewDecoder(resp.Body).Decode(&urls); err != nil {
        return nil, err
    }

    return urls, nil
}

func (wc *WorkerCrawler) Crawl() {
    urls, err := wc.FetchURLs()
    if err != nil {
        log.Fatalf("Failed to fetch URLs: %v", err)
    }

    for _, urlObj := range urls {
        wc.crawlURL(urlObj.Address)
    }
}

func (wc *WorkerCrawler) crawlURL(startURL string) {
    collector := colly.NewCollector()

    baseURL, err := url.Parse(startURL)
    if err != nil {
        log.Printf("Invalid start URL: %v", err)
        return
    }

    collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        if strings.Contains(link, "#") {
            log.Printf("Ignoring link with fragment: %s", link)
            return
        }

        linkURL, err := url.Parse(link)
        if err != nil {
            log.Printf("Invalid link URL: %v", err)
            return
        }

        absoluteURL := baseURL.ResolveReference(linkURL).String()
        log.Printf("Found link: %s", absoluteURL)
    })

    collector.Visit(startURL)
}