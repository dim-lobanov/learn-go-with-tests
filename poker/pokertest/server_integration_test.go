package pokertest

import (
	"fmt"
	"learn-go-with-tests/poker"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]") // [] - empty valid json
	defer cleanDatabase()
	// store := NewInMemoryPlayerStore()
	store, err := poker.NewFileSystemPlayerStore(database)
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}

	server, err := poker.NewPlayerServer(store, &GameSpy{})
	if err != nil {
		t.Fatal("problem creating player server", err)
	}

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
		want := []poker.Player{{Name: "Pepper", Wins: 3}}

		assertLeague(t, got, want)
	})
}
