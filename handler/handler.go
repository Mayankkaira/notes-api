package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Mayankkaira/notes-api/application"
)

type Handler struct {
	application *application.Application
}

func NewHandler(application *application.Application) *Handler {
	handler := &Handler{application: application}
	return handler
}

type CreateNoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, errVal ErrorResponse, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(errVal)
	if err != nil {
		log.Printf("failed to encode error response: %v", err)
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := CreateNoteRequest{}
	err := decoder.Decode(&request)
	if err != nil {
		errorResponse := ErrorResponse{Error: "invalid JSON"}
		WriteErrorResponse(w, errorResponse, 400)
		return
	}
	createNote, err := h.application.CreateNote(request.Title, request.Content)
	if err != nil {
		if errors.Is(err, application.ErrInvalidInput) {
			createNoteErr := ErrorResponse{Error: err.Error()}
			WriteErrorResponse(w, createNoteErr, 400)
			return
		} else {
			createNoteErr := ErrorResponse{Error: "internal server error"}
			WriteErrorResponse(w, createNoteErr, 500)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encodeErr := encoder.Encode(createNote)
	if encodeErr != nil {
		log.Printf("failed to encode note response %v", encodeErr)
	}

}
