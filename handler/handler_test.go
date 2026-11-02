package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/note"
	"github.com/Mayankkaira/notes-api/storage"
)

func TestHandler_Create(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	title := "Learn Go"
	content := "Practice testing"
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
	recorder := httptest.NewRecorder()
	h.Create(recorder, req)
	if recorder.Code != http.StatusCreated {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusCreated)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	decoder := json.NewDecoder(recorder.Body)
	rNote := note.Note{}
	err := decoder.Decode(&rNote)
	if err != nil {
		t.Fatalf("create decoding error:%v", err)
	}
	if rNote.Title != title {
		t.Errorf("title:%v want Learn Go", rNote.Title)
	}
	if rNote.Content != content {
		t.Errorf("Content:%v want Practice testing", rNote.Content)
	}
}

func TestHandler_Create_Invalid_JSON(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":`))
	recorder := httptest.NewRecorder()
	h.Create(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	decoder := json.NewDecoder(recorder.Body)
	eNote := ErrorResponse{}
	err := decoder.Decode(&eNote)
	if err != nil {
		t.Fatalf("create decoding error:%v", err)
	}
	if eNote.Error != "invalid JSON" {
		t.Errorf("create() error: %v want invalid JSON", eNote.Error)
	}
}

func TestHandler_Create_Invalid_Input(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"","content":"Practice testing"`))
	recorder := httptest.NewRecorder()
	h.Create(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	decoder := json.NewDecoder(recorder.Body)
	eNote := ErrorResponse{}
	err := decoder.Decode(&eNote)
	if err != nil {
		t.Fatalf("create decoding error:%v", err)
	}
	if recorder.Code != http.StatusBadRequest {
		t.Errorf("status code : %d, want %d", recorder.Code, http.StatusBadRequest)
	}
}
