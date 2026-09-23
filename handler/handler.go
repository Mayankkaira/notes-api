package handler

import (
	"encoding/json"
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
		errorResponse := ErrorResponse{Error: err.Error()}
		WriteErrorResponse(w, errorResponse, 400)
		return
	}
	h.application.CreateNote(request.Title, request.Content)

}
