package database

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

// ConnectPostgres establishes a connection to PostgreSQL
func ConnectPostgres(url string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
