package service

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/google/uuid"
)

type imageRepository interface {
	AddImage(image domain.Image) (int, error)
	GetImages(userID int) ([]domain.Image, error)
	UploadImage(ctx context.Context, bucketName, objectName string, reader io.Reader) (*url.URL, error)
}

// ImageService represents a service layer for image.
type ImageService struct {
	repo imageRepository
}

// NewImageService creates and returns a new ImageService object
func NewImageService(repo imageRepository) *ImageService {
	return &ImageService{repo: repo}
}

// AddImage insert image into repository
func (s *ImageService) AddImage(image domain.Image) (int, error) {
	return s.repo.AddImage(image)
}

// GetImages retrieves all images owned by certain user
func (s *ImageService) GetImages(userID int) ([]domain.Image, error) {
	return s.repo.GetImages(userID)
}

func generateUniqueFilename(filename string) string {
	id := uuid.New()
	ext := filepath.Ext(filename)
	return fmt.Sprintf("%s_%s%s", id.String(), strings.TrimSuffix(filename, ext), ext)
}

// UploadImage uploads image to the specified bucket
func (s *ImageService) UploadImage(ctx context.Context, bucketName, objectName string, reader io.Reader) (*url.URL, error) {
	return s.repo.UploadImage(ctx, bucketName, generateUniqueFilename(objectName), reader)
}
