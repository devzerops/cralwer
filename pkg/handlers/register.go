package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterCommonHandlers(r *mux.Router) {
	r.Handle("/info", InfoHandler()).Methods("GET")
	r.Handle("/info/ip", InfoIPHandler()).Methods("GET")
	r.Handle("/config", ConfigHandler()).Methods("GET")
	r.Handle("/config", ConfigHandler()).Methods("POST")
	r.Handle("/health", HealthHandler()).Methods("GET")
	r.Handle("/status", StatusHandler()).Methods("GET")
}