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
	storedNode, err := a.storage.Get(id)
	return storedNode, err
}

func (a *Application) ListNotes() ([]note.Note, error) {
	cNotes, err := a.storage.List()
	return cNotes, err
}

func (a *Application) UpdateNote(id int, title, content *string) error {
	if title == nil && content == nil {
		return fmt.Errorf("both fields can't be empty")
	}
	if title != nil {
		if len(*title) == 0 {
			return fmt.Errorf("Invalid title")
		}
	}
	if content != nil {
		if len(*content) == 0 {
			return fmt.Errorf("Invalid content")
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
	return err
}
