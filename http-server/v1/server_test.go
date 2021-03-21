package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseBody(t, res.Body.String(), "20")
	})

	t.Run("return Floyed's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)
		assertResponseBody(t, res.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t testing.TB, receive, expect string) {
	t.Helper()
	if receive != expect {
		t.Errorf("got %q, want %q", receive, expect)
	}
}
