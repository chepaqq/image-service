package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/chepaqq/image-service/internal/delivery/api/handler"
	customMiddleware "github.com/chepaqq/image-service/internal/delivery/api/middleware"
)

// NewRouter initializes routes using Chi
func NewRouter(userHandler *handler.UserHandler, imageHandler *handler.ImageHandler, jwtSecret string) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Public routes
	router.Post("/login", userHandler.SignIn)
	router.Post("/register", userHandler.SignUp)

	// Protected routes
	router.Route("/api", func(r chi.Router) {
		r.Use(customMiddleware.AuthMiddleware(jwtSecret))
		r.Use(customMiddleware.AccessControlMiddleware)

		r.Get("/images", imageHandler.GetImages)
		r.Post("/upload-picture", imageHandler.UploadImage)
	})

	return router
}
