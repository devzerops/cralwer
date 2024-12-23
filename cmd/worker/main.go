package main

import (
	commonHandlers "distributed-crawler/m/pkg/handlers"
	workerHandlers "distributed-crawler/m/pkg/worker/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	registerWorkerHandlers(r)
	registerCommonHandlers(r)


	log.Fatal(http.ListenAndServe(":8081", r))
}

func registerWorkerHandlers(r *mux.Router) {
	r.HandleFunc("/download", workerHandlers.DownloadHandler).Methods("GET")
	r.HandleFunc("/parse", workerHandlers.ParseHandler).Methods("GET")
}

func registerCommonHandlers(r *mux.Router) {
	r.HandleFunc("/status", commonHandlers.StatusHandler).Methods("GET")
	r.HandleFunc("/config", commonHandlers.ConfigHandler).Methods("GET")
	r.HandleFunc("/config", commonHandlers.ConfigHandler).Methods("POST")
	r.HandleFunc("/health", commonHandlers.HealthHandler).Methods("GET")
}
