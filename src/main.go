package main

import (
	"log"
	"net/http"
)

func main() {
	// -- middleware handling in controller.go
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":80", router))
}