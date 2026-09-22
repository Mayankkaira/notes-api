package handler

import "github.com/Mayankkaira/notes-api/application"

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
