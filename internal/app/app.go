package app

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chepaqq/jungle-task/internal/config"
	"github.com/chepaqq/jungle-task/internal/delivery/api/handler"
	"github.com/chepaqq/jungle-task/internal/delivery/api/middleware"
	"github.com/chepaqq/jungle-task/internal/repository"
	"github.com/chepaqq/jungle-task/internal/service"
	"github.com/chepaqq/jungle-task/pkg/database"
	"github.com/chepaqq/jungle-task/pkg/storage"
	"github.com/gorilla/mux"
)

// Run initialize and starts application
func Run(cfg *config.Config) {
	// Deps
	postgresURL := fmt.Sprintf(
		"user=%s dbname=%s host=%s password=%s port=%s sslmode=%s",
		cfg.Postgres.User,
		cfg.Postgres.DBName,
		cfg.Postgres.Host,
		cfg.Postgres.Password,
		cfg.Postgres.Port,
		cfg.Postgres.SSLMode,
	)
	postgresClient, err := database.ConnectPostgres(postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	useSSL, err := strconv.ParseBool(cfg.Minio.SSL)
	if err != nil {
		log.Fatal(err)
	}

	minioClient, err := storage.ConnectMinio(useSSL, cfg.Minio.Endpoint, cfg.Minio.BucketName, cfg.Minio.BucketLocation)
	if err != nil {
		log.Fatal(err)
	}

	// Repos
	userRepository := repository.NewUserRepository(postgresClient)
	imageRepository := repository.NewImageRepository(postgresClient, minioClient)

	// Services
	userService := service.NewUserService(userRepository)
	imageService := service.NewImageService(imageRepository)

	// Middlewares
	userMiddleware := middleware.NewUserMiddleware(*userService)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	imageHandler := handler.NewImageHandler(imageService)

	// Routes
	router := mux.NewRouter()

	router.HandleFunc("/login", userHandler.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/register", userHandler.SignUp).Methods(http.MethodPost)

	restrictRouter := router.PathPrefix("/").Subrouter()
	restrictRouter.Use(userMiddleware.AccessMiddleware)
	restrictRouter.HandleFunc("/images", imageHandler.GetImages).Methods(http.MethodGet)
	restrictRouter.HandleFunc("/upload-picture", imageHandler.UploadImage).Methods(http.MethodPost)

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
