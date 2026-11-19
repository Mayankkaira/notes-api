package storage

import "database/sql"

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) *PostgresStorage {
	newStorage := &PostgresStorage{db: db}
	return newStorage
}
