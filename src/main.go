package main

import (
	"log"
	"net/http"
	"middleware"
)

func main() {
	// set up routes
	router := NewRouter()
	// middleware
	router.Use(middleware.Authentication)
	router.Use(middleware.HttpHeaders)
	// start server
	log.Fatal(http.ListenAndServe(":80", router))
}