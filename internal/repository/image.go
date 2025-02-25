package repository

import (
	"context"
	"io"
	"net/url"

	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/chepaqq/jungle-task/pkg/storage"

	"github.com/jmoiron/sqlx"
)

// ImageRepository represents repository for an image entity
type ImageRepository struct {
	db      *sqlx.DB
	storage storage.Storage
}

// NewImageRepository creates and returns a new ImageRepository object
func NewImageRepository(db *sqlx.DB, storage storage.Storage) *ImageRepository {
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

// GetImages retrieves from repository all images owned by certain user
func (r *ImageRepository) GetImages(userID int) ([]domain.Image, error) {
	var images []domain.Image
	query := `SELECT * FROM image WHERE user_id=$1`
	err := r.db.Select(&images, query, userID)
	if err != nil {
		return nil, err
	}
	return images, nil
}

// UploadImage uploads object to Minio bucket
func (r *ImageRepository) UploadImage(ctx context.Context, bucketName, objectName string, reader io.Reader) (*url.URL, error) {
	return r.storage.Upload(ctx, bucketName, objectName, reader, "image/jpeg")
}
