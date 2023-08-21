package repository

import (
	"context"
	"io"
	"net/url"
	"path/filepath"

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
	_, err := r.storage.PutObject(ctx, bucketName, objectName, reader, -1, minio.PutObjectOptions{ContentType: "image/jpeg"})
	if err != nil {
		return nil, err
	}
	url := r.storage.EndpointURL()
	// TODO: Retrieve from env file
	url.Host = "0.0.0.0:9000"
	url.Path = filepath.Join(url.Path, "images", objectName)
	return url, nil
}
