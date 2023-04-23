package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]") // [] - empty valid json
	defer cleanDatabase()
	// store := NewInMemoryPlayerStore()
	store, err := NewFileSystemPlayerStore(database)
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}

	server := NewPlayerServer(store)
	const player = "Pepper"
	const playersWins = 3

	for i := 0; i < playersWins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), fmt.Sprint(playersWins))
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		got := getLeagueFromResponse(t, response.Body)
		want := []Player{{"Pepper", 3}}

		assertLeague(t, got, want)
	})
}
