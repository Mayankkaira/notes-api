package application

import (
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
	Title   *string
	Content *string
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

func NewApplication(storage Storage)*Application{
          app:=&Application{storage:storage}
		  return app
}