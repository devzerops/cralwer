package models

type CrawlRequest struct {
	URL     string `json:"url"`
	Depth   int    `json:"depth"`
	Timeout int    `json:"timeout"`
}

type CrawlResult struct {
	URL   string   `json:"url"`
	Title string   `json:"title"`
	Links []string `json:"links"`
	Error string   `json:"error,omitempty"`
}
