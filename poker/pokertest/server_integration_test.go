package pokertest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"learn-go-with-tests/poker"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := createTempFile(t, "[]") // [] - empty valid json
	defer cleanDatabase()
	// store := NewInMemoryPlayerStore()
	store, err := poker.NewFileSystemPlayerStore(database)
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}

	server := poker.NewPlayerServer(store)
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
