package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {

	t.Run("returns Pepper's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		if err != nil {
			t.Errorf("Something wrong happened %w", err)
		}

		res := httptest.NewRecorder()

		PlayerServer(res, req)

		received := res.Body.String()
		expected := "20"

		if received != expected {
			t.Errorf("got %q, want %q", received, expected)
		}
	})

	t.Run("return Floyed's score", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		if err != nil {
			t.Errorf("Something wrong happened %w", err)
		}

		res := httptest.NewRecorder()

		PlayerServer(res, req)

		received := res.Body.String()
		expected := "10"

		if received != expected {
			t.Errorf("got %q, want %q", received, expected)
		}
	})
}
