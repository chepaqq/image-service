package handler

import (
	"encoding/json"
	"net/http"

	"github.com/chepaqq/jungle-task/internal/domain"
)

type userService interface {
	CreateUser(user domain.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

// UserHandler handles HTTP requests related to user
type UserHandler struct {
	userService userService
}

// NewUserHandler creates and returns a new UserHandler object
func NewUserHandler(userService userService) *UserHandler {
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
	id, err := h.userService.CreateUser(input)
	// TODO: check if user exists
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// SignIn handles user authorization
func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input signInInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input body", http.StatusBadRequest)
		return
	}
	token, err := h.userService.GenerateToken(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
