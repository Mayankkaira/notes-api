package main

import (
	"fmt"
	"net/http"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/handler"
	"github.com/Mayankkaira/notes-api/storage"
)

func main() {
	newStorage := storage.NewInMemoryStorage()
	app := application.NewApplication(newStorage)
	newHandler := handler.NewHandler(app)
	http.HandlerFunc("POST/notes", newHandler.Create)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
