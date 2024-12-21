package main

import (
	"distributed-crawler/server"
	"log"
	"net/http"
)

func main() {
	serverRouter := server.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", serverRouter))
	
	// workerRouter := worker.NewRouter()
	// log.Fatal(http.ListenAndServe(":8081", workerRouter))
}
