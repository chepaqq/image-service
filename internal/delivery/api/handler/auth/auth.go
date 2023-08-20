package auth

import (
	"encoding/json"
	"net/http"

	"github.com/chepaqq/jungle-task/internal/domain"
	"github.com/chepaqq/jungle-task/internal/service/auth"
)

type Handler struct {
	authService auth.Service
}

func NewHandler(authService auth.Service) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	var input domain.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input body", http.StatusBadRequest)
		return
	}
	id, err := h.authService.CreateUser(input)
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
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	var input domain.User
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "invalid input body", http.StatusBadRequest)
		return
	}
	token, err := h.authService.GenerateToken(input.Username, input.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
	})
}
