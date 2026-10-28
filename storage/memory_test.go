package storage

import (
	"errors"
	"testing"
	"time"

	"github.com/Mayankkaira/notes-api/application"
)

func TestInMemoryStorage_Create(t *testing.T) {
	s := NewInMemoryStorage()
	currentTime := time.Now()
	input := application.CreateNoteInput{
		Title:     "Learn Go",
		Content:   "Practise Test",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	note, err := s.Create(input)
	if err != nil {
		t.Fatalf("create note %v", err)
	}
	if note.ID != 1 {
		t.Error("first id should be 1")
	}
	rNote, err := s.Get(note.ID)
	if err != nil {
		t.Fatalf("get note:%v", err)
	}
	if rNote.ID != note.ID {
		t.Errorf("retirved note id =%v want %v", rNote.ID, note.ID)
	}
	if rNote.Title != input.Title {
		t.Errorf("%v doesn't match with %v", rNote.Title, input.Title)
	}
	if rNote.Content != input.Content {
		t.Errorf("%v doesn't match with %v", rNote.Content, input.Content)
	}

}

func TestInMemoryStorage_Create_IncrementID(t *testing.T) {
	s := NewInMemoryStorage()
	currentTime := time.Now()
	input1 := application.CreateNoteInput{
		Title:     "Learn Go",
		Content:   "Practise Test",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	input2 := application.CreateNoteInput{
		Title:     "Learn python",
		Content:   "Class and objects",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	note1, err1 := s.Create(input1)
	note2, err2 := s.Create(input2)
	if err1 != nil {
		t.Fatalf("create note %v", err1)
	}
	if err2 != nil {
		t.Fatalf("create note %v", err2)
	}
	if note1.ID != 1 || note2.ID != 2 {
		t.Errorf("ID assignment incorrect: first ID = %d (want 1), second ID = %d (want 2)", note1.ID, note2.ID)
	}
	rNote, err := s.Get(note2.ID)
	if err != nil {
		t.Fatalf("get note:%v", err)
	}
	if rNote.Title != input2.Title {
		t.Errorf("%v doesn't match with %v", rNote.Title, input2.Title)
	}
	if rNote.Content != input2.Content {
		t.Errorf("%v doesn't match with %v", rNote.Content, input2.Content)
	}
}

func TestInMemory_Update_TitleOnly(t *testing.T) {
	s := NewInMemoryStorage()
	currentTime := time.Now()
	input := application.CreateNoteInput{
		Title:     "Learn Go",
		Content:   "Practise Testing",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	rNote, err := s.Create(input)
	if err != nil {
		t.Fatalf("Create note %v", err)
	}
	newTitle := "Learn concurrency"
	updatedInput := application.UpdateNoteInput{
		Title:     &newTitle,
		UpdatedAt: time.Now(),
	}
	uNote, err := s.Update(rNote.ID, updatedInput)
	if err != nil {
		t.Fatalf("update note %v", err)
	}
	if uNote.Title != newTitle {
		t.Errorf("updated Title: %q want %q", uNote.Title, newTitle)
	}
	if uNote.Content != rNote.Content {
		t.Errorf("updated content = %q, want %q", uNote.Content, rNote.Content)
	}
	if !uNote.UpdatedAt.Equal(updatedInput.UpdatedAt) {
		t.Errorf("updated time = %v, want %v", uNote.UpdatedAt, updatedInput.UpdatedAt)
	}
}

func TestInMemory_Update_ContentOnly(t *testing.T) {
	s := NewInMemoryStorage()
	currentTime := time.Now()
	input := application.CreateNoteInput{
		Title:     "Learn Go",
		Content:   "Practise Testing",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	rNote, err := s.Create(input)
	if err != nil {
		t.Fatalf("Create note %v", err)
	}
	newContent := "testing update"
	updatedInput := application.UpdateNoteInput{
		Content:   &newContent,
		UpdatedAt: time.Now(),
	}
	uNote, err := s.Update(rNote.ID, updatedInput)
	if err != nil {
		t.Fatalf("update note %v", err)
	}
	if uNote.Content != newContent {
		t.Errorf("updated content: %q want %q", uNote.Content, newContent)
	}
	if uNote.Title != rNote.Title {
		t.Errorf("updated Title = %q, want %q", uNote.Title, rNote.Title)
	}
	if !uNote.UpdatedAt.Equal(updatedInput.UpdatedAt) {
		t.Errorf("updated time = %v, want %v", uNote.UpdatedAt, updatedInput.UpdatedAt)
	}
}

func TestInMemoryStorage_Update_Nofound(t *testing.T) {
	s := NewInMemoryStorage()
	newTitle := "Learn go"
	updatedInput := application.UpdateNoteInput{
		Title:     &newTitle,
		UpdatedAt: time.Now(),
	}
	_, err := s.Update(2, updatedInput)
	if !errors.Is(err, application.ErrStorageNotFound) {
		t.Fatalf("Update() error = %v, want %v", err, application.ErrStorageNotFound)
	}
}

func TestInMemoryStorage_Delete(t *testing.T) {
	s := NewInMemoryStorage()
	currentTime := time.Now()
	input := application.CreateNoteInput{
		Title:     "Learn Go",
		Content:   "Practise Test",
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	rNote, err := s.Create(input)
	if err != nil {
		t.Fatalf("create note:%v", err)
	}
	err = s.Delete(rNote.ID)
	if err != nil {
		t.Fatalf("Delete() error: %v", err)
	}
	_, err = s.Get(rNote.ID)
	if !errors.Is(err, application.ErrStorageNotFound) {
		t.Errorf("error= %v, want %v", err, application.ErrStorageNotFound)
	}
}

func TestInMemoryStorage_Delete_NotFound(t *testing.T) {
	s := NewInMemoryStorage()
	err := s.Delete(2)
	if !errors.Is(err, application.ErrStorageNotFound) {
		t.Fatalf("delete() %v, want %v", err, application.ErrStorageNotFound)
	}
}
