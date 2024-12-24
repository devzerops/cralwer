package handlers

import (
	"net/http"
	"distributed-crawler/m/pkg/utils"
)

func StatusHandler() http.Handler {
	return utils.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Status OK"))
	}))
}