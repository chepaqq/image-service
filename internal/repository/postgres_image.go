package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/chepaqq/image-service/internal/domain"
)

type PostgresImageRepository struct {
	db *sqlx.DB
}

// NewImageRepository creates and returns a new ImageRepository object
func NewImageRepository(db *sqlx.DB) *PostgresImageRepository {
	return &PostgresImageRepository{db: db}
}

// AddImage inserts new image to the repository
func (r *PostgresImageRepository) AddImage(image domain.Image) (int, error) {
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
func (r *PostgresImageRepository) GetImages(userID int) ([]domain.Image, error) {
	var images []domain.Image
	query := `SELECT * FROM image WHERE user_id=$1`
	err := r.db.Select(&images, query, userID)
	if err != nil {
		return nil, err
	}
	return images, nil
}
