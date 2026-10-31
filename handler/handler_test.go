package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/storage"
)

func TestHandler_Create(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
	recorder := httptest.NewRecorder()
	h.Create(recorder, req)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("something went wrong status code want %d", http.StatusCreated)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
}
