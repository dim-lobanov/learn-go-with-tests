package main

// To create a web server in Go you will typically call ListenAndServe:
// func ListenAndServe(addr string, handler Handler) error

// type Handler interface { ServeHTTP(ResponseWriter, *Request) }

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type PlayerServer struct {
	store  PlayerStore
	http.Handler // embedding https://go.dev/doc/effective_go#embedding
}

// we don't want to set up router every call of ServeHttp
// We are implementing Handler interface
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux() // ServeMux is an HTTP request multiplexer
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	// What this means is that our PlayerServer now has all the methods that http.Handler has, which is just ServeHTTP.
	// To "fill in" the http.Handler we assign it to the router we create in NewPlayerServer. 
	// We can do this because http.ServeMux has the method ServeHTTP.
	p.Handler = router

	return p
}

const jsonContentType = "application/json"

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

// You can use embedding with interfaces to compose new interface
//
// type Animal interface {
// 	  Eater
// 	  Sleeper
// }
// As you'd expect if you embed a concrete type you'll have access to all its public methods and fields.
// When embedding types, really think about what impact that has on your public API.
