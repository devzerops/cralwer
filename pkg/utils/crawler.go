package utils

import (
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"github.com/segmentio/kafka-go"
)

// Crawler struct holds the Colly collector and Kafka writer
type Crawler struct {
	collector   *colly.Collector
	kafkaWriter *kafka.Writer
	keywords    []string
}

// NewCrawler initializes a new Crawler with Colly and Kafka configurations
func NewCrawler(kafkaBroker string, keywords []string) *Crawler {
    c := colly.NewCollector(
        colly.Async(true),
        colly.MaxDepth(3),
    )

    // Limit the crawling rate to avoid hitting the same domain too frequently
    c.Limit(&colly.LimitRule{
        DomainGlob:  "*",
        Parallelism: 2,
        Delay:       5 * time.Second,
    })

    kafkaWriter := &kafka.Writer{
        Addr:     kafka.TCP(kafkaBroker),
        Topic:    "crawled_urls",
        Balancer: &kafka.LeastBytes{},
    }

    return &Crawler{
        collector:   c,
        kafkaWriter: kafkaWriter,
        keywords:    keywords,
    }
}

// Start begins the crawling process for the given URLs
func (c *Crawler) Start(urls []string) {
	for _, u := range urls {
		c.collector.Visit(u)
	}
	c.collector.Wait()
}

// setupHandlers sets up the handlers for the Colly collector
func (c *Crawler) setupHandlers() {
    // Handle links found in the crawled pages
    c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
        link := e.Attr("href")
        parsedURL, err := url.Parse(link)
        if err == nil && parsedURL.Scheme != "" && parsedURL.Host != "" {
            if IsMaliciousURLByKeyword(parsedURL.String(), c.keywords) {
                e.Request.Abort()
            } else {
                c.kafkaWriter.WriteMessages(nil, kafka.Message{
                    Value: []byte(parsedURL.String()),
                })
            }
        }
    })

    // Handle images found in the crawled pages
    c.collector.OnHTML("img[src]", func(e *colly.HTMLElement) {
        imgSrc := e.Attr("src")
        parsedURL, err := url.Parse(imgSrc)
        if err == nil && parsedURL.Scheme != "" && parsedURL.Host != "" {
            if IsMaliciousURLByKeyword(parsedURL.String(), c.keywords) {
                e.Request.Abort()
            } else {
                c.kafkaWriter.WriteMessages(nil, kafka.Message{
                    Value: []byte(parsedURL.String()),
                })
            }
        }
    })

    // Abort requests to malicious URLs
    c.collector.OnRequest(func(r *colly.Request) {
        if IsMaliciousURLByKeyword(r.URL.String(), c.keywords) {
            r.Abort()
        }
    })
}

// isMaliciousURL checks if the URL contains any malicious keywords
func isMaliciousURL(u string) bool {
	maliciousKeywords := []string{"malware", "phishing", "scam", "crime", "violence"}
	for _, keyword := range maliciousKeywords {
		if strings.Contains(u, keyword) {
			return true
		}
	}
	return false
}
