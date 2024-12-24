package main

import (
	commonHandlers "distributed-crawler/m/pkg/handlers"
	workerHandlers "distributed-crawler/m/pkg/worker/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	workerHandlers.RegisterWorkerHandlers(r)
	commonHandlers.RegisterCommonHandlers(r)

	log.Fatal(http.ListenAndServe(":8081", r))
}