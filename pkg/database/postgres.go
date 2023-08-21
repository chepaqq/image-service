package database

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// ConnectPostgres establishes a connection to PostgreSQL
func ConnectPostgres(url string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", url)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
