package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	commonHandlers "distributed-crawler/m/pkg/handlers"
)

func main() {
	r := mux.NewRouter()

	// Register handlers
	commonHandlers.RegisterCommonHandlers(r)

	log.Fatal(http.ListenAndServe(":8082", r))
}
