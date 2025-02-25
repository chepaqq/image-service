package service

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/chepaqq/image-service/internal/domain"
	"github.com/chepaqq/image-service/internal/repository"
	"github.com/chepaqq/image-service/pkg/storage"

	"github.com/google/uuid"
)

type ImageService interface {
	AddImage(image domain.Image) (int, error)
	GetImages(userID int) ([]domain.Image, error)
	UploadImage(ctx context.Context, bucketName, objectName string, reader io.Reader) (*url.URL, error)
}

type imageService struct {
	repo    repository.ImageRepository
	storage storage.Storage
}

// NewImageService creates and returns a new imageService object
func NewImageService(repo repository.ImageRepository, storage storage.Storage) ImageService {
	return &imageService{repo: repo, storage: storage}
}

// AddImage insert image into repository
func (s *imageService) AddImage(image domain.Image) (int, error) {
	return s.repo.AddImage(image)
}

// GetImages retrieves all images owned by certain user
func (s *imageService) GetImages(userID int) ([]domain.Image, error) {
	return s.repo.GetImages(userID)
}

// UploadImage uploads image to the specified bucket
func (s *imageService) UploadImage(ctx context.Context, bucketName, objectName string, reader io.Reader) (*url.URL, error) {
	id := uuid.New()
	ext := filepath.Ext(objectName)
	objectName = fmt.Sprintf("%s_%s%s", id.String(), strings.TrimSuffix(objectName, ext), ext)

	return s.storage.Upload(ctx, bucketName, objectName, reader, "image/jpeg")
}
