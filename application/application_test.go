package application

import (
	"errors"
	"testing"

	"github.com/Mayankkaira/notes-api/note"
)

type FakeStorage struct {
	createCalled bool
}

func (f *FakeStorage) Create(input CreateNoteInput) (note.Note, error) {
	f.createCalled = true
	return note.Note{Title: input.Title, Content: input.Content}, nil
}

func (f *FakeStorage) Delete(id int) error {
	return nil
}

func (f *FakeStorage) Get(id int) (note.Note, error) {
	return note.Note{}, nil
}
func (f *FakeStorage) List() ([]note.Note, error) {
	return []note.Note{}, nil
}
func (f *FakeStorage) Update(ID int, updateInput UpdateNoteInput) (note.Note, error) {
	return note.Note{}, nil
}

func TestCreateNote_EmptyTitle(t *testing.T) {
	fake := &FakeStorage{}
	app := NewApplication(fake)
	title := ""
	content := "something"
	_, err := app.CreateNote(title, content)
	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("%v is not %v", err, ErrInvalidInput)
	}
	if fake.createCalled {
		t.Errorf("storage.Create() was called for invalid input")
	}
}

func TestCreateNote_EmptyContent(t *testing.T) {
	fake := &FakeStorage{}
	app := NewApplication(fake)
	title := "Learn Go"
	content := ""
	_, err := app.CreateNote(title, content)
	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("%v is not %v", err, ErrInvalidInput)
	}
	if fake.createCalled {
		t.Errorf("storage.Create() was called for invalid input")
	}
}
