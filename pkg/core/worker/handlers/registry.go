package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterWorkerHandlers(r *mux.Router) {
    r.HandleFunc("/download", DownloadHandler).Methods("GET")
    r.HandleFunc("/parse", ParseHandler).Methods("GET")
}