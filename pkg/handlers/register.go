package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterCommonHandlers(r *mux.Router) {
	r.HandleFunc("/status", StatusHandler).Methods("GET")
	r.HandleFunc("/config", ConfigHandler).Methods("GET")
	r.HandleFunc("/config", ConfigHandler).Methods("POST")
	r.HandleFunc("/health", HealthHandler).Methods("GET")
}