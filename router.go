package main

import (
	"net/http"

	"github.com/Mayankkaira/notes-api/handler"
)

func NewRouter(h *handler.Handler) *http.ServeMux {
	newMux := http.NewServeMux()
	newMux.HandleFunc("POST /notes", h.Create)
	newMux.HandleFunc("GET /notes", h.List)
	newMux.HandleFunc("GET /notes/{id}", h.Get)
	newMux.HandleFunc("PATCH /notes/{id}", h.Update)
	newMux.HandleFunc("DELETE /notes/{id}", h.Delete)
	return newMux
}
