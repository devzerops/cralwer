package handlers

import (
	"encoding/json"
	"net/http"
	"distributed-crawler/m/pkg/utils"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler() http.Handler {
	return utils.LoggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		health := HealthResponse{
			Status: "Healthy",
		}

		response, err := json.Marshal(health)
		if (err != nil) {
			http.Error(w, "Failed to marshal health response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}))
}