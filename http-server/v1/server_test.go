package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	t.Run("returns Pepper's score", func(t *testing.T) {
		req := newGetScoreRequest("Pepper")
		res := httptest.NewRecorder()

		PlayerServer(res, req)
		assertResponseBody(t, res.Body.String(), "20")
	})

	t.Run("return Floyed's score", func(t *testing.T) {
		req := newGetScoreRequest("Floyd")
		res := httptest.NewRecorder()

		PlayerServer(res, req)
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
