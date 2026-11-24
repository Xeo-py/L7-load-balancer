package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	targetURL, err := url.Parse("http://localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)
	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", proxy))
}
