package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	PlayerServer(res, req)

	t.Run("returns Pepper's score", func(t *testing.T) {
		received := res.Body.String()
		expected := "20"

		if received != expected {
			t.Errorf("got %q, want %q", received, expected)
		}
	})

}
