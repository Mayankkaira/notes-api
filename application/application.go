package application

import (
	"fmt"
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
	Create(input CreateNoteInput) (int, error)
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

func (a *Application) CreateNote(title, content string) (int, error) {
	if len(title) == 0 || len(content) == 0 {
		return 0, fmt.Errorf("empty value error")
	}
	currentTime := time.Now()
	noteInput := CreateNoteInput{
		Title:     title,
		Content:   content,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	id, err := a.storage.Create(noteInput)
	return id, err
}

func (a *Application) GetNote(id int) (note.Note, error) {
	storedNote, err := a.storage.Get(id)
	return storedNote, err
}
