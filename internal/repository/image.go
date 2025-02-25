package repository

import (
	"github.com/chepaqq/image-service/internal/domain"
)

type ImageRepository interface {
	AddImage(image domain.Image) (int, error)
	GetImages(userID int) ([]domain.Image, error)
}
