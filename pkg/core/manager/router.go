package manager

import (
	"net/http"
	"github.com/gorilla/mux"
)
type Manager struct {
	port string
}

func (m *Manager) NewManager(port string) *Manager {
	router := mux.NewRouter()
	router.HandleFunc("/api/start", m.startHandler).Methods("POST")
	router.HandleFunc("/api/stop", m.stopHandler).Methods("POST")
	http.ListenAndServe(":"+m.port, router)
	return m
}

func (m *Manager) startHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to start the necessary processes
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Started"))
}

func (m *Manager) stopHandler(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to stop the necessary processes
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stopped"))
}