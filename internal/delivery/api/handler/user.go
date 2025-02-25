package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/chepaqq/image-service/internal/domain"
	"github.com/chepaqq/image-service/internal/service"
	"github.com/chepaqq/image-service/pkg/logger"
)

// UserHandler handles HTTP requests related to user
type UserHandler struct {
	userService service.UserService
}

// NewUserHandler creates and returns a new UserHandler object
func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// SignUp handles user registration
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input domain.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input body", http.StatusBadRequest)
		return
	}
	if len(input.Password) < 8 {
		http.Error(w, "short password", http.StatusBadRequest)
		logger.Error("short password")
		return
	}
	id, err := h.userService.CreateUser(input)
	if err != nil {
		if errors.Is(err, domain.ErrUserConflict) {
			http.Error(w, "user already exists", http.StatusConflict)
			logger.Error("user already exists")
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// SignIn handles user authorization
func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input body", http.StatusBadRequest)
		logger.Error("invalid input body")
		return
	}

	if input.Username == "" || input.Password == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		logger.Error("missing required fields")
		return
	}

	if len(input.Password) < 8 {
		http.Error(w, "short password", http.StatusBadRequest)
		logger.Error("short password")
		return
	}
	token, err := h.userService.GenerateToken(input.Username, input.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			logger.Error("invalid username or password")
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
