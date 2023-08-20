package auth

import (
	"net/http"

	"github.com/chepaqq/jungle-task/internal/service/auth"
	"github.com/gorilla/mux"
)

type Handler struct {
	authService auth.Service
}

func NewHandler(authService auth.Service) *Handler {
	return &Handler{authService: authService}
}

func (h *Handler) InitRoutes(r *mux.Router) {
}

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
}
