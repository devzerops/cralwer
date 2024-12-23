package worker

import (
	"distributed-crawler/m/pkg/models"
)

type Worker struct {
	port string
}

func NewWorker(port string) *Worker {
	return &Worker{port: port}
}

func (w *Worker) Download(url string) ([]byte, error) {
	// HTML 다운로드 로직 구현
	return nil, nil
}

func (w *Worker) Parse(html []byte) (*models.CrawlResult, error) {
	// HTML 파싱 로직 구현
	return nil, nil
}
