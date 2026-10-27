package storage

import (
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
