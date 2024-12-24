package handlers

import (
	"distributed-crawler/m/pkg/models"
	"github.com/gocolly/colly"
	"log"
)

type Worker struct{
	*models.Worker
	Logger *log.Logger
}

type WorkerConfig struct{
	*models.WorkerConfig
}

func (w *Worker) DownloadHTML(url string) (string, error) {
	c := colly.NewCollector()

	var htmlContent string

	c.OnHTML("html", func(e *colly.HTMLElement) {
		var err error
		htmlContent, err = e.DOM.Html()
		if err != nil {
			w.Logger.Println("Error getting HTML content:", err)
		}
	})

	c.OnError(func(_ *colly.Response, err error) {
		w.Logger.Println("Something went wrong:", err)
	})

	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	return htmlContent, nil
}