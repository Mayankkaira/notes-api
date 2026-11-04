package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/handler"
	"github.com/Mayankkaira/notes-api/note"
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
	t.Run("List", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		app := application.NewApplication(s)
		h := handler.NewHandler(app)
		router := NewRouter(h)
		recorder := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/notes", nil)
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusOK {
			t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusOK)
		}
		if recorder.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
		}
		notes := []note.Note{}
		err := json.NewDecoder(recorder.Body).Decode(&notes)
		if err != nil {
			t.Fatalf("List() decoder error: %v", err)
		}
		if notes == nil {
			t.Fatalf("List is nill want empty[]")
		}
		if len(notes) != 0 {
			t.Fatalf("len(notes) = %d, want 0", len(notes))
		}
	})
	t.Run("Get", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		app := application.NewApplication(s)
		h := handler.NewHandler(app)
		router := NewRouter(h)
		recorder := httptest.NewRecorder()
		title := "Learn Go"
		content := "Practice testing"
		req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusCreated {
			t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusCreated)
		}
		if recorder.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
		}
		recorder1 := httptest.NewRecorder()
		req1 := httptest.NewRequest(http.MethodGet, "/notes/1", nil)
		router.ServeHTTP(recorder1, req1)
		if recorder1.Code != http.StatusOK {
			t.Fatalf("status code = %d, want %d", recorder1.Code, http.StatusOK)
		}
		if recorder1.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder1.Header().Get("Content-Type"))
		}
		notes := note.Note{}
		decoder := json.NewDecoder(recorder1.Body)
		err := decoder.Decode(&notes)
		if err != nil {
			t.Fatalf("Get() decoder error: %v", err)
		}
		if notes.ID != 1 {
			t.Fatalf("Get() id: %d, want %d", notes.ID, 1)
		}
		if notes.Title != title {
			t.Errorf("title: %q want %q", notes.Title, title)
		}
		if notes.Content != content {
			t.Errorf("Content: %q want %q", notes.Content, content)
		}
	})
	t.Run("Update", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		app := application.NewApplication(s)
		h := handler.NewHandler(app)
		router := NewRouter(h)
		req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusCreated {
			t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusCreated)
		}
		if recorder.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
		}
		req1 := httptest.NewRequest(http.MethodPatch, "/notes/1", strings.NewReader(`{"title":"Learn Advance Go","content":"Practice concurrency"}`))
		recorder1 := httptest.NewRecorder()
		router.ServeHTTP(recorder1, req1)
		if recorder1.Code != http.StatusOK {
			t.Fatalf("status code = %d, want %d", recorder1.Code, http.StatusOK)
		}
		if recorder1.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder1.Header().Get("Content-Type"))
		}
		decoder := json.NewDecoder(recorder1.Body)
		rNote := note.Note{}
		err := decoder.Decode(&rNote)
		if err != nil {
			t.Fatalf("Update() decoding error:%v", err)
		}
		if rNote.ID != 1 {
			t.Fatalf("note id: %d want %d", rNote.ID, 1)
		}
		if rNote.Title != "Learn Advance Go" {
			t.Errorf("title: %q, want %q", rNote.Title, "Learn Advance Go")
		}
		if rNote.Content != "Practice concurrency" {
			t.Errorf("content: %q, want %q", rNote.Content, "Practice concurrency")
		}
	})
	t.Run("Delete", func(t *testing.T) {
		s := storage.NewInMemoryStorage()
		app := application.NewApplication(s)
		h := handler.NewHandler(app)
		router := NewRouter(h)
		req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)
		if recorder.Code != http.StatusCreated {
			t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusCreated)
		}
		if recorder.Header().Get("Content-Type") != "application/json" {
			t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
		}
		req1 := httptest.NewRequest(http.MethodDelete, "/notes/1", nil)
		recorder1 := httptest.NewRecorder()
		router.ServeHTTP(recorder1, req1)
		if recorder1.Code != http.StatusNoContent {
			t.Fatalf("status code = %d, want %d", recorder1.Code, http.StatusNoContent)
		}
		if recorder1.Body.Len() != 0 {
			t.Errorf("response body = %q, want empty body", recorder1.Body.String())
		}
		req2 := httptest.NewRequest(http.MethodGet, "/notes/1", nil)
		recorder2 := httptest.NewRecorder()
		router.ServeHTTP(recorder2, req2)
		if recorder2.Code != http.StatusNotFound {
			t.Fatalf("status code = %d, want %d", recorder2.Code, http.StatusNotFound)
		}
	})
}
