package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chepaqq/image-service/internal/delivery/api/handler"
	"github.com/chepaqq/image-service/internal/delivery/api/middleware"
)

// NewRouter initializes routes
func NewRouter(userHandler handler.UserHandler, imageHandler handler.ImageHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", userHandler.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/register", userHandler.SignUp).Methods(http.MethodPost)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.ApplicationRecovery)

	restrictRouter := router.PathPrefix("/").Subrouter()
	restrictRouter.HandleFunc("/images", imageHandler.GetImages).Methods(http.MethodGet)
	restrictRouter.HandleFunc("/upload-picture", imageHandler.UploadImage).Methods(http.MethodPost)

	restrictRouter.Use(middleware.AuthMiddleware)
	restrictRouter.Use(middleware.AccessControlMiddleware)
	return router
}
