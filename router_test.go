package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/handler"
	"github.com/Mayankkaira/notes-api/storage"
)

func TestRouter_Routes(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		app := application.NewApplication(s)
		h := handler.NewHandler(app)
		router := NewRouter(h)
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusCreated {
			t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusCreated)
		}
		if recorder.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
		}
	})
}
