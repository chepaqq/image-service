package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chepaqq/jungle-task/internal/delivery/api/middleware"
	"github.com/chepaqq/jungle-task/internal/domain"
)

type imageService interface {
	GetImages(userID int) ([]domain.Image, error)
	UploadImage(ctx context.Context, bucketName, objectName string, reader io.Reader) (*url.URL, error)
	AddImage(image domain.Image) (int, error)
}

// ImageHandler handles HTTP requests related to image
type ImageHandler struct {
	imageService imageService
}

// NewImageHandler creates and returns a new ImageHandler object
func NewImageHandler(imageService imageService) *ImageHandler {
	return &ImageHandler{imageService: imageService}
}

// UploadImage handles HTTP request to upload image
func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "user_id not found in context", http.StatusInternalServerError)
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	src, hdr, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		return
	}
	defer src.Close()

	url, err := h.imageService.UploadImage(r.Context(), "images", hdr.Filename, src)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	h.imageService.AddImage(domain.Image{
		UserID:    userID,
		ImageURL:  url.String(),
		ImagePath: hdr.Filename,
	})
	json.NewEncoder(w).Encode(map[string]interface{}{
		"url": url.String(),
	})
}

// GetImages handles HTTP request to retrieve all images
func (h *ImageHandler) GetImages(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "user_id not found in context", http.StatusInternalServerError)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	images, err := h.imageService.GetImages(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(images)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
