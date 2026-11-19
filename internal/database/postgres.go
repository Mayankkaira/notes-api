package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Mayankkaira/notes-api/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func OpenDB(ctx context.Context, cfg config.Config) (*sql.DB, error) {
	connectionString := cfg.DSN()
	db, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("ping database %w", err)
	}
	return db, nil
}
