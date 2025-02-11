package main

import (
	"log"
	"net/http"
)

var port = ":8080"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", HandleWebsocket)

	go HandleBroadcast()

	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Printf("Server started at port %v\n", port)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
