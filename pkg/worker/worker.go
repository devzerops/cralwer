package worker

import (
	"distributed-crawler/m/pkg/models"
	"distributed-crawler/m/pkg/storage"
	"log"
)

type Worker struct {
	models.Worker
}

func NewWorker(port, processID, ip string, cassandraClusterIPs []string, cassandraKeyspace string) *Worker {
	// Cassandra 초기화
	err := storage.InitCassandra(cassandraClusterIPs, cassandraKeyspace)
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	worker := &Worker{
		Worker: models.Worker{
			Port:      port,
			ProcessID: processID,
			IP:        ip,
		},
	}
	err = storage.UpdateProcessInfo(processID, ip, "running")
	if err != nil {
		log.Printf("Failed to update process info: %v", err)
	}
	return worker
}

func (w *Worker) Close() {
	err := storage.DeleteProcessInfo(w.ProcessID)
	if err != nil {
		log.Printf("Failed to delete process info: %v", err)
	}
	storage.CloseCassandra()
}

func (w *Worker) Download(url string) ([]byte, error) {
	// HTML 다운로드 로직 구현
	return nil, nil
}

func (w *Worker) Parse(html []byte) (*models.CrawlResult, error) {
	// HTML 파싱 로직 구현
	return nil, nil
}
