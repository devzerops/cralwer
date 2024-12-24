package handlers

import (
	"github.com/gorilla/mux"
)

func RegisterServerHandlers(r *mux.Router) {
	r.HandleFunc("/collect", CollectHandler).Methods("POST")
	r.HandleFunc("/resolve", ResolveHandler).Methods("POST")
}
