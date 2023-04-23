package main

import (
	"log"
	"net/http"
	"os"
	"learn-go-with-tests/poker"
)

const dbFileName = "game.db.json"

func main() {
	// The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.
	// If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.
	// HandlerFunc implements ServeHTTP method
	//
	// handler := http.HandlerFunc(PlayerServerFunc) // btw it's a type conversion

	database, err := os.OpenFile(dbFileName, os.O_RDWR | os.O_CREATE, 0666) // RDWR - ReaD, WRite; chmod 666
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemPlayerStore(database)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}
	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server)) // takes a port to listen to, starts server (http.Handler) basically
	// ListenAndServe returns error, so we can wrap it with log
}
