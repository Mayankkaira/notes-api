package storage

import (
	"fmt"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/note"
)

type InMemoryStorage struct {
	notes  map[int]note.Note
	nextID int
}

func NewInMemoryStorage() *InMemoryStorage {
	notes := make(map[int]note.Note)
	return &InMemoryStorage{notes: notes, nextID: 1}
}

func (s *InMemoryStorage) Create(input application.CreateNoteInput) (note.Note, error) {
	unote := note.Note{
		Title:     input.Title,
		Content:   input.Content,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}
	unote.ID = s.nextID
	s.notes[unote.ID] = unote
	s.nextID++
	return unote, nil
}

func (s *InMemoryStorage) Get(id int) (note.Note, error) {
	val, ok := s.notes[id]
	if !ok {
		return val, application.ErrStorageNotFound
	}
	return val, nil
}

func (s *InMemoryStorage) List() ([]note.Note, error) {
	var noteSlice []note.Note
	for _, value := range s.notes {
		noteSlice = append(noteSlice, value)
	}
	return noteSlice, nil
}

func (s *InMemoryStorage) Update(id int, update application.UpdateNoteInput) error {
	val, ok := s.notes[id]
	if !ok {
		return fmt.Errorf("Note with %d doesn't exist", id)
	}
	if update.Title != nil {
		val.Title = *update.Title
	}
	if update.Content != nil {
		val.Content = *update.Content
	}
	val.UpdatedAt = update.UpdatedAt
	s.notes[id] = val
	return nil
}

func (s *InMemoryStorage) Delete(id int) error {
	_, ok := s.notes[id]
	if !ok {
		return application.ErrStorageNotFound
	}
	delete(s.notes, id)
	return nil
}
