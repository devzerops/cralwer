package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	commonHandlers "distributed-crawler/m/pkg/handlers"
	serverHandlers "distributed-crawler/m/pkg/server/handlers"
)

func main() {
	r := mux.NewRouter()

	// Register handlers
	serverHandlers.RegisterServerHandlers(r)
	commonHandlers.RegisterCommonHandlers(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}