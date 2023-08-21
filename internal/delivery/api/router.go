package api

import (
	"net/http"

	"github.com/chepaqq/jungle-task/internal/delivery/api/handler"
	"github.com/chepaqq/jungle-task/internal/delivery/api/middleware"
	"github.com/gorilla/mux"
)

// NewRouter initializes routes
func NewRouter(userHandler handler.UserHandler, imageHandler handler.ImageHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", userHandler.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/register", userHandler.SignUp).Methods(http.MethodPost)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.LoggingMiddleware)

	restrictRouter := router.PathPrefix("/").Subrouter()
	restrictRouter.HandleFunc("/images", imageHandler.GetImages).Methods(http.MethodGet)
	restrictRouter.HandleFunc("/upload-picture", imageHandler.UploadImage).Methods(http.MethodPost)

	restrictRouter.Use(middleware.AuthMiddleware)
	restrictRouter.Use(middleware.AccessControlMiddleware)
	return router
}
