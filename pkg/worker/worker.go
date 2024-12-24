package worker

import (
	"distributed-crawler/m/pkg/models"
	"distributed-crawler/m/pkg/storage"
	"log"
	"time"
)

type WorkerConfig struct {
	Port             string
	ProcessID        string
	IP               string
	CassandraIP      string
	CassandraKeyspace string
}

type Worker struct {
	models.Worker
	stopHeartbeat chan bool
}

func NewWorker(config WorkerConfig) *Worker {
	// Cassandra 초기화
	err := storage.InitCassandra(config.CassandraIP, config.CassandraKeyspace)
	if (err != nil) {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	worker := &Worker{
		Worker: models.Worker{
			Port:      config.Port,
			ProcessID: config.ProcessID,
			IP:        config.IP,
		},
		stopHeartbeat: make(chan bool),
	}
	err = storage.UpdateProcessInfo(config.ProcessID, config.IP, "running")
	if err != nil {
		log.Printf("Failed to update process info: %v", err)
	}

	go worker.startHeartbeat()

	return worker
}

func (w *Worker) startHeartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := storage.UpdateProcessInfo(w.ProcessID, w.IP, "running")
			if err != nil {
				log.Printf("Failed to update heartbeat: %v", err)
			}
		case <-w.stopHeartbeat:
			return
		}
	}
}

func (w *Worker) Close() {
	w.stopHeartbeat <- true
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
