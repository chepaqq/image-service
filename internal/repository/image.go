package repository

import (
	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/minio/minio-go/v7"
)

// ImageRepository represents repository for an image entity
type ImageRepository struct {
	db      *sqlx.DB
	storage *minio.Client
}

// NewImageRepository creates and returns a new ImageRepository object
func NewImageRepository(db *sqlx.DB, storage *minio.Client) *ImageRepository {
	return &ImageRepository{db: db, storage: storage}
}

// AddImage inserts new image to the repository
func (r *ImageRepository) AddImage(image domain.Image) (int, error) {
	var id int
	query := `INSERT INTO image(user_id, image_path, image_url) VALUES ($1, $2, $3)`
	row := r.db.QueryRow(query, image.UserID, image.ImagePath, image.ImageURL)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
