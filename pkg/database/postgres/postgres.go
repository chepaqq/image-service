package db

import (
	"fmt"
	"time"

	"github.com/chepaqq99/jungle-task/internal/config"
	"github.com/jmoiron/sqlx"
)

func ConnectPostgres(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open(
		"postgres",
		fmt.Sprintf("user=%s dbname=%s host=%s password=%s port=%s sslmode=%s", cfg.User, cfg.DBName, cfg.Host, cfg.Password, cfg.Port, cfg.SSLMode),
	)
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
