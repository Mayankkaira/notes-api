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
		t.Errorf("title: %q want %q", rNote.Title, title)
	}
	if rNote.Content != content {
		t.Errorf("Content: %q want %q", rNote.Content, content)
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
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"","content":"Practice testing"}`))
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
	if eNote.Error != application.ErrInvalidInput.Error() {
		t.Errorf("invalid input error: %v, want %v", eNote.Error, application.ErrInvalidInput.Error())
	}
}

func TestHandler_Create_EmptyContent(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn go","content":""}`))
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
	if eNote.Error != application.ErrInvalidInput.Error() {
		t.Errorf("invalid input error: %v, want %v", eNote.Error, application.ErrInvalidInput.Error())
	}
}

func TestHandler_List_Empty(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	recorder := httptest.NewRecorder()
	h.List(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusOK)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	decoder := json.NewDecoder(recorder.Body)
	rNote := []note.Note{}
	err := decoder.Decode(&rNote)
	if err != nil {
		t.Fatalf("List decoding error:%v", err)
	}
	if len(rNote) != 0 {
		t.Fatalf("List returned: %d, want 0", len(rNote))
	}
	if rNote == nil {
		t.Fatalf("List() returned nil slice, want non-nil empty slice")
	}
}

func TestHandler_List_Populated(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req1 := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn Go","content":"Practice testing"}`))
	recorder1 := httptest.NewRecorder()
	h.Create(recorder1, req1)
	if recorder1.Code != http.StatusCreated {
		t.Fatalf("status code = %d, want %d", recorder1.Code, http.StatusCreated)
	}
	if recorder1.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder1.Header().Get("Content-Type"))
	}
	req2 := httptest.NewRequest(http.MethodPost, "/notes", strings.NewReader(`{"title":"Learn HTTP","content":"Build REST APIs"}`))
	recorder2 := httptest.NewRecorder()
	h.Create(recorder2, req2)
	if recorder2.Code != http.StatusCreated {
		t.Fatalf("status code = %d, want %d", recorder2.Code, http.StatusCreated)
	}
	if recorder2.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder2.Header().Get("Content-Type"))
	}
	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	recorder := httptest.NewRecorder()
	h.List(recorder, req)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusOK)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	rNote := []note.Note{}
	decoder := json.NewDecoder(recorder.Body)
	err := decoder.Decode(&rNote)
	if err != nil {
		t.Fatalf("list() decoder error: %v", err)
	}
	if len(rNote) != 2 {
		t.Fatalf("list() returned: %d want 2", len(rNote))
	}
	foundNote1 := false
	foundNote2 := false
	for _, value := range rNote {
		if value.ID == 1 {
			foundNote1 = true
			if value.Title != "Learn Go" {
				t.Errorf("List() returned %v want %q", value.Title, "Learn Go")
			}
			if value.Content != "Practice testing" {
				t.Errorf("List() returned %v want %q", value.Content, "Practice testing")
			}
		}
		if value.ID == 2 {
			foundNote2 = true
			if value.Title != "Learn HTTP" {
				t.Errorf("List()returned %q want %q", value.Title, "Learn HTTP")
			}
			if value.Content != "Build REST APIs" {
				t.Errorf("List()returned %q want %q", value.Content, "Build REST APIs")
			}
		}
	}
	if !foundNote1 {
		t.Error("created note1 not found")
	}
	if !foundNote2 {
		t.Error("created note2 not found")
	}
}

func TestHandler_Get_Success(t *testing.T) {
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
	req1 := httptest.NewRequest(http.MethodGet, "/notes/1", nil)
	req1.SetPathValue("id", "1")
	recorder1 := httptest.NewRecorder()
	h.Get(recorder1, req1)
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
		t.Fatalf("Get decoding error:%v", err)
	}
	if rNote.ID != 1 {
		t.Fatalf("Get() id: %d, want %d", rNote.ID, 1)
	}
	if rNote.Title != title {
		t.Errorf("title: %q want %q", rNote.Title, title)
	}
	if rNote.Content != content {
		t.Errorf("Content: %q want %q", rNote.Content, content)
	}
}

func TestHandler_Get_NotFound(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodGet, "/notes/1", nil)
	req.SetPathValue("id", "1")
	recorder := httptest.NewRecorder()
	h.Get(recorder, req)
	if recorder.Code != http.StatusNotFound {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusNotFound)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	decoder := json.NewDecoder(recorder.Body)
	errNote := ErrorResponse{}
	err := decoder.Decode(&errNote)
	if err != nil {
		t.Fatalf("Get decoding error:%v", err)
	}
	if errNote.Error != application.ErrNoteNotFound.Error() {
		t.Errorf("Get(): %v, want %v", errNote.Error, application.ErrNoteNotFound.Error())
	}
}

func TestHandler_Get_InvalidID(t *testing.T) {
	s := storage.NewInMemoryStorage()
	app := application.NewApplication(s)
	h := NewHandler(app)
	req := httptest.NewRequest(http.MethodGet, "/notes/abc", nil)
	req.SetPathValue("id", "abc")
	recorder := httptest.NewRecorder()
	h.Get(recorder, req)
	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("status code = %d, want %d", recorder.Code, http.StatusBadRequest)
	}
	if recorder.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("Content-Type = %v,want application/json ", recorder.Header().Get("Content-Type"))
	}
	decoder := json.NewDecoder(recorder.Body)
	errNote := ErrorResponse{}
	err := decoder.Decode(&errNote)
	if err != nil {
		t.Fatalf("Get decoding error:%v", err)
	}
	if errNote.Error != "invalid id" {
		t.Errorf("Get(): %q, want %q", errNote.Error, "invalid id")
	}
}
