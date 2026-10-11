package main

import (
	"log"
	"net/http"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/handler"
	"github.com/Mayankkaira/notes-api/storage"
)

func main() {
	newStorage := storage.NewInMemoryStorage()
	app := application.NewApplication(newStorage)
	newHandler := handler.NewHandler(app)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /notes", newHandler.Create)
	mux.HandleFunc("GET /notes", newHandler.List)
	mux.HandleFunc("GET /notes/{id}", newHandler.Get)
	mux.HandleFunc("PATCH /notes/{id}", newHandler.Update)
	mux.HandleFunc("DELETE /notes/{id}", newHandler.Delete)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
