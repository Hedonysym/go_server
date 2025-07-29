package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", readyEndpointHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
