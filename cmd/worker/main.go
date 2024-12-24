package main

import (
	"distributed-crawler/m/pkg/worker"
	commonHandlers "distributed-crawler/m/pkg/handlers"
	workerHandlers "distributed-crawler/m/pkg/worker/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	// Get port, process ID, IP, Cassandra IP, and keyspace from environment variables.
	config := worker.WorkerConfig{
		Port:             os.Getenv("WORKER_PORT"),
		ProcessID:        os.Getenv("PROCESS_ID"),
		IP:               os.Getenv("WORKER_IP"),
		CassandraIP:      os.Getenv("CASSANDRA_IP"),
		CassandraKeyspace: os.Getenv("CASSANDRA_KEYSPACE"),
	}

	// Create worker
	w := worker.NewWorker(config)
	defer w.Close()

	// Register handlers
	workerHandlers.RegisterWorkerHandlers(r)
	commonHandlers.RegisterCommonHandlers(r)

	// Start server
	go func() {
		log.Fatal(http.ListenAndServe(":"+config.Port, r))
	}()

	// Handle shutdown signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down worker...")
}