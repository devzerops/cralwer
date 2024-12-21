package crawler

import (
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly/v2"
)

func CrawlURL(startURL string, onLinkFound func(string)) {
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
		onLinkFound(absoluteURL)
	})

	collector.Visit(startURL)
}
