package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/chepaqq/image-service/internal/delivery/api/middleware"
	"github.com/chepaqq/image-service/internal/domain"
	"github.com/chepaqq/image-service/internal/service"
	"github.com/chepaqq/image-service/pkg/logger"
)

// ImageHandler handles HTTP requests related to image
type ImageHandler struct {
	imageService service.ImageService
}

// NewImageHandler creates and returns a new ImageHandler object
func NewImageHandler(imageService service.ImageService) *ImageHandler {
	return &ImageHandler{imageService: imageService}
}

// UploadImage handles HTTP request to upload image
func (h *ImageHandler) UploadImage(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User_id not found in context", http.StatusInternalServerError)
		logger.Error("user_id not found in context")
		return
	}
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Failed to convert string to int", http.StatusInternalServerError)
		logger.Error("failed to convert string to int")
	}

	err = r.ParseMultipartForm(10 << 20) // 10 MB max file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		logger.Error("unable to parse form")
		return
	}

	src, hdr, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusBadRequest)
		logger.Error("Unable to get file")
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

	w.Header().Set("Content-Type", "application/json")
	if images != nil {
		err = json.NewEncoder(w).Encode(images)
	} else {
		err = json.NewEncoder(w).Encode("User has no images")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
