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
	mux := NewRouter(newHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
