package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "8081", "Port to run backend server on")
	flag.Parse()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Served from backend on port %s\n", *port)
		log.Printf("Received request from %s\n", r.RemoteAddr)
	})
	address := ":" + *port
	log.Printf("Backend server listening on http://localhost%s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Server failed to start:%v", err)
	}
}
