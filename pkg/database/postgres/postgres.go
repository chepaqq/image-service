package db

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func ConnectPostgres(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("user=%s dbname=%s host=%s password=%s port=%s sslmode=%s", cfg.Username, cfg.DBName, cfg.Host, cfg.Password, cfg.Port, cfg.SSLMode),
	)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	if err != nil {
		return nil, err
	}
	return db, nil
}
