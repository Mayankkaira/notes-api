package database

import (
	"context"
	"database/sql"
	"fmt"
    "time"
	"github.com/Mayankkaira/notes-api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenDB(ctx context.Context, cfg config.Config) (*sql.DB, error) {
	connectionString := cfg.DSN()
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(2 * time.Minute)
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database %w", err)
	}
	return db, nil
}
