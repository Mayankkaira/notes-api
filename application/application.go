package application

import (
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
	Create(input CreateNoteInput) (note.Note, error)
	Get(ID int) (note.Note, error)
	List() ([]note.Note, error)
	Update(ID int, updateInput UpdateNoteInput) error
	Delete(ID int) error
}

type Application struct {
	storage Storage
}

func NewApplication(storage Storage) *Application {
	app := &Application{storage: storage}
	return app
}

func (a *Application) CreateNote(title, content string) (note.Note, error) {
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
	createdNote, err := a.storage.Create(noteInput)
	return createdNote, err
}

func (a *Application) GetNote(id int) (note.Note, error) {
	storedNote, err := a.storage.Get(id)
	if errors.Is(err, ErrStorageNotFound) {
		return storedNote, ErrNoteNotFound
	}
	return storedNote, err
}

func (a *Application) ListNotes() ([]note.Note, error) {
	cNotes, err := a.storage.List()
	return cNotes, err
}

func (a *Application) UpdateNote(id int, title, content *string) error {
	if title == nil && content == nil {
		return ErrInvalidInput
	}
	if title != nil {
		if len(*title) == 0 {
			return ErrInvalidInput
		}
	}
	if content != nil {
		if len(*content) == 0 {
			return ErrInvalidInput
		}
	}
	updateNote := UpdateNoteInput{
		Title:     title,
		Content:   content,
		UpdatedAt: time.Now(),
	}
	err := a.storage.Update(id, updateNote)
	return err

}

func (a *Application) DeleteNote(id int) error {
	err := a.storage.Delete(id)
	if errors.Is(err, ErrStorageNotFound) {
		return ErrNoteNotFound
	}
	return err
}
