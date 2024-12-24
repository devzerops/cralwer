package handlers

import (
	"github.com/gocolly/colly"
	"net/http"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	c := colly.NewCollector()
	c.OnResponse(func(response *colly.Response) {
		w.Write(response.Body)
	})

	err := c.Visit(url)
	if err != nil {
		http.Error(w, "Failed to download the page", http.StatusInternalServerError)
	}
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	c := colly.NewCollector()
	c.OnHTML("title", func(e *colly.HTMLElement) {
		w.Write([]byte(e.Text))
	})

	err := c.Visit(url)
	if err != nil {
		http.Error(w, "Failed to parse the page", http.StatusInternalServerError)
	}
}


