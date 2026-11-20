package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mayankkaira/notes-api/application"
	"github.com/Mayankkaira/notes-api/handler"
	"github.com/Mayankkaira/notes-api/internal/config"
	"github.com/Mayankkaira/notes-api/internal/database"
	"github.com/Mayankkaira/notes-api/internal/storage"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("configuration failure %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := database.OpenDB(ctx, cfg)
	if err != nil {
		log.Fatalf("database connection failed %v", err)
	}
	defer db.Close()
	postgresStorage:=storage.NewPostgresStorage(db)
	app := application.NewApplication(postgresStorage)
	newHandler := handler.NewHandler(app)
	mux := NewRouter(newHandler)
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
