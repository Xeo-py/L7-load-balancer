package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 1. Define a terminal flag named "port" with a default value of "8081"
	port := flag.String("port", "8081", "Port to run the backend server on")
	flag.Parse()

	// 2. Define the HTTP handler for the root path "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Print a message back to the user's browser or curl command
		fmt.Fprintf(w, "Served from backend on port: %s\n", *port)

		// Log the incoming request to the terminal so you can see it happen
		log.Printf("Received request from %s\n", r.RemoteAddr)
	})

	// 3. Format the address (e.g., ":8081") and start listening
	address := ":" + *port
	log.Printf("Backend server listening on http://localhost%s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
