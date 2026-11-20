package application

import (
	"context"
	"errors"
	"time"

	"github.com/Mayankkaira/notes-api/note"
)

type CreateNoteInput struct {
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UpdateNoteInput struct {
	Title     *string
	Content   *string
	UpdatedAt time.Time
}

type Storage interface {
	Create(ctx context.Context, input CreateNoteInput) (note.Note, error)
	Get(ctx context.Context, ID int) (note.Note, error)
	List(ctx context.Context) ([]note.Note, error)
	Update(ctx context.Context, ID int, updateInput UpdateNoteInput) (note.Note, error)
	Delete(ctx context.Context, ID int) error
}

type Application struct {
	storage Storage
}

func NewApplication(storage Storage) *Application {
	app := &Application{storage: storage}
	return app
}

func (a *Application) CreateNote(ctx context.Context, title, content string) (note.Note, error) {
	var createdNote note.Note
	if len(title) == 0 || len(content) == 0 {
		return createdNote, ErrInvalidInput
	}
	currentTime := time.Now()
	noteInput := CreateNoteInput{
		Title:     title,
		Content:   content,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	createdNote, err := a.storage.Create(ctx, noteInput)
	return createdNote, err
}

func (a *Application) GetNote(ctx context.Context, id int) (note.Note, error) {
	storedNote, err := a.storage.Get(ctx, id)
	if errors.Is(err, ErrStorageNotFound) {
		return storedNote, ErrNoteNotFound
	}
	return storedNote, err
}

func (a *Application) ListNotes(ctx context.Context) ([]note.Note, error) {
	cNotes, err := a.storage.List(ctx)
	return cNotes, err
}

func (a *Application) UpdateNote(ctx context.Context, id int, title, content *string) (note.Note, error) {
	var updatedNote note.Note
	if title == nil && content == nil {
		return updatedNote, ErrInvalidInput
	}
	if title != nil {
		if len(*title) == 0 {
			return updatedNote, ErrInvalidInput
		}
	}
	if content != nil {
		if len(*content) == 0 {
			return updatedNote, ErrInvalidInput
		}
	}
	updateNote := UpdateNoteInput{
		Title:     title,
		Content:   content,
		UpdatedAt: time.Now(),
	}
	updatedNote, err := a.storage.Update(ctx, id, updateNote)
	if errors.Is(err, ErrStorageNotFound) {
		return updatedNote, ErrNoteNotFound
	}
	return updatedNote, err

}

func (a *Application) DeleteNote(ctx context.Context, id int) error {
	err := a.storage.Delete(ctx, id)
	if errors.Is(err, ErrStorageNotFound) {
		return ErrNoteNotFound
	}
	return err
}
