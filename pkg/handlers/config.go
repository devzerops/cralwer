package handlers

import (
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Config OK"))
}