package handler

import "github.com/Mayankkaira/notes-api/application"

type Handler struct {
	application *application.Application
}

func NewHandler(application *application.Application) *Handler {
	handler := &Handler{application: application}
	return handler
}
