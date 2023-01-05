package racer_test

import (
	"net/http"
	"net/http/httptest"
	"select/racer"
	"testing"
	"time"
)

func TestRacerTestRacer(t *testing.T) {
	t.Run("correctly receives the faster server response", func(t *testing.T) {
		slowServer := makeDelayedServer(5 * time.Millisecond)
		fastServer := makeDelayedServer(1 * time.Millisecond)
		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := racer.Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		slowServer.Close()
		fastServer.Close()
	})

	t.Run("returns an error when no server responds within specified time", func(t *testing.T) {
		serverA := makeDelayedServer(12 * time.Millisecond)
		serverB := makeDelayedServer(15 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := racer.ConfigurableRacer(serverA.URL, serverB.URL, 5*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
