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

type Process struct {
	ID       string   `json:"id"`
	Type	 string   `json:"type"`
	Platform string   `json:"platform"`
}