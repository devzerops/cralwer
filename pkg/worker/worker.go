package worker

import (
	"distributed-crawler/m/pkg/models"
	"distributed-crawler/m/pkg/storage/cassandra"
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
	Port      string
	ProcessID string
	IP        string
	stopHeartbeat chan bool
}

func NewWorker(config WorkerConfig) *Worker {
	// Initialize Cassandra
	err := cassandra.InitCassandra(config.CassandraIP, config.CassandraKeyspace)
	if err != nil {
		log.Fatalf("Failed to initialize Cassandra: %v", err)
	}

	worker := &Worker{
		Port:          config.Port,
		ProcessID:     config.ProcessID,
		IP:            config.IP,
		stopHeartbeat: make(chan bool),
	}
	err = cassandra.UpdateProcessInfo(config.ProcessID, config.IP, "running")
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
			err := cassandra.UpdateProcessInfo(w.ProcessID, w.IP, "running")
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
	err := cassandra.DeleteProcessInfo(w.ProcessID)
	if err != nil {
		log.Printf("Failed to delete process info: %v", err)
	}
	cassandra.CloseCassandra()
}

func (w *Worker) Download(url string) ([]byte, error) {
	// Implement HTML download logic
	return nil, nil
}

func (w *Worker) Parse(html []byte) (*models.CrawlResult, error) {
	// Implement HTML parsing logic
	return nil, nil
}
