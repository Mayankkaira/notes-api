package application

import (
	"context"
	"errors"
	"testing"

	"github.com/Mayankkaira/notes-api/note"
)

type FakeStorage struct {
	createCalled bool
}

func (f *FakeStorage) Create(ctx context.Context, input CreateNoteInput) (note.Note, error) {
	f.createCalled = true
	return note.Note{Title: input.Title, Content: input.Content, CreatedAt: input.CreatedAt, UpdatedAt: input.UpdatedAt}, nil
}

func (f *FakeStorage) Delete(ctx context.Context, id int) error {
	return nil
}

func (f *FakeStorage) Get(ctx context.Context, id int) (note.Note, error) {
	return note.Note{}, nil
}
func (f *FakeStorage) List(ctx context.Context) ([]note.Note, error) {
	return []note.Note{}, nil
}
func (f *FakeStorage) Update(ctx context.Context, ID int, updateInput UpdateNoteInput) (note.Note, error) {
	return note.Note{}, nil
}

func TestCreateNote_EmptyTitle(t *testing.T) {
	fake := &FakeStorage{}
	app := NewApplication(fake)
	title := ""
	content := "something"
	_, err := app.CreateNote(t.Context(), title, content)
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
	_, err := app.CreateNote(t.Context(), title, content)
	if !errors.Is(err, ErrInvalidInput) {
		t.Errorf("%v is not %v", err, ErrInvalidInput)
	}
	if fake.createCalled {
		t.Errorf("storage.Create() was called for invalid input")
	}
}

func TestCreateNote_Success(t *testing.T) {
	fake := &FakeStorage{}
	app := NewApplication(fake)
	title := "Learn Go"
	content := "Practice Testing"
	note, err := app.CreateNote(t.Context(), title, content)
	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}
	if !fake.createCalled {
		t.Error("storage.Create() is not called")
	}
	if note.Title != title {
		t.Errorf("created title %v doesn't match with input title:%v", note.Title, title)
	}
	if note.Content != content {
		t.Errorf("created content %v doesgn't match with input content:%v", note.Content, content)
	}
	if note.CreatedAt.IsZero() {
		t.Error("Note Created time should not be zero ")
	}
	if note.UpdatedAt.IsZero() {
		t.Error("Note Updated time should not be zero")
	}
	if !note.CreatedAt.Equal(note.UpdatedAt) {
		t.Error("created and updated time should be same")
	}
}
