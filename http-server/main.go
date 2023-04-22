package main

import (
	"log"
	"net/http"
)

func main() {
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	// HandlerFunc implements ServeHTTP method
	//
	// handler := http.HandlerFunc(PlayerServerFunc) // btw it's a type conversion

	server := NewPlayerServer(NewInMemoryPlayerStore())
	log.Fatal(http.ListenAndServe(":5000", server)) // takes a port to listen to
	// ListenAndServe returns error, so we can wrap it with log
}
