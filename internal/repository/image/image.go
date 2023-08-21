package auth

import (
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

// Repository holds a database and storage connections
type Repository struct {
	db      *sqlx.DB
	storage *minio.Client
}

// NewRepository creates and returns a new Repository object
func NewRepository(db *sqlx.DB, storage *minio.Client) *Repository {
	return &Repository{db: db, storage: storage}
}
