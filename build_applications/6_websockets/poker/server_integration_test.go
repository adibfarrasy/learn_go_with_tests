package poker_test

import (
	"net/http"
	"net/http/httptest"
	. "server/poker"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	database, cleanDatabase := CreateTempFile(t, "")
	defer cleanDatabase()
	store, err := NewFileSystemPlayerStore(database)
	AssertNoError(t, err)

	server, _ := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), NewPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewGetScoreRequest(player))
		AssertStatus(t, response, http.StatusOK)

		AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, NewLeagueRequest())
		AssertStatus(t, response, http.StatusOK)

		got := GetLeagueFromResponse(t, response.Body)
		want := []Player{
			{Name: "Pepper", Wins: 3},
		}
		AssertLeague(t, got, want)
	})
}
