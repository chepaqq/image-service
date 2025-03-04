package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/chepaqq/image-service/internal/config"
	"github.com/chepaqq/image-service/internal/delivery/api"
	"github.com/chepaqq/image-service/internal/delivery/api/handler"
	"github.com/chepaqq/image-service/internal/repository"
	"github.com/chepaqq/image-service/internal/service"
	"github.com/chepaqq/image-service/pkg/database"
	"github.com/chepaqq/image-service/pkg/logger"
	"github.com/chepaqq/image-service/pkg/server"
	"github.com/chepaqq/image-service/pkg/storage"
)

// Run initialize and starts application
func Run(cfg *config.Config) {
	// Deps
	postgresURL := fmt.Sprintf(
		"user=%s dbname=%s host=%s password=%s port=%s sslmode=%s",
		cfg.Database.User,
		cfg.Database.DBName,
		cfg.Database.Host,
		cfg.Database.Password,
		cfg.Database.Port,
		cfg.Database.SSLMode,
	)
	postgresClient, err := database.ConnectPostgres(postgresURL)
	if err != nil {
		logger.Fatalf("Failed to connect to Postgres: %v", err)
	}
	logger.Info("Starting Postgres")

	minioStorage, err := storage.NewMinioStorage(cfg.Storage.Endpoint, cfg.Storage.BucketName, cfg.Storage.BucketLocation, cfg.Storage.SSL)
	if err != nil {
		logger.Fatalf("Failed to connect to Minio: %v", err)
	}

	logger.Info("Starting Minio")

	// Repos
	userRepository := repository.NewPostgresUserRepository(postgresClient)
	imageRepository := repository.NewPostgresImageRepository(postgresClient)

	// Services
	userService := service.NewUserService(cfg, userRepository)
	imageService := service.NewImageService(imageRepository, minioStorage)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	imageHandler := handler.NewImageHandler(imageService)

	// HTTP
	router := api.NewRouter(userHandler, imageHandler, cfg.Auth.JWTSecret)
	server := server.New(router, cfg.Server.Port)

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-interrupt:
		logger.Errorf("Signal interrupt error: %s", s.String())
	case err = <-server.Notify():
		logger.Infof("Server notify %v:", err)
	}

	// Shutdown server
	err = server.Shutdown()
	if err != nil {
		logger.Infof("Server shutdown: %v", err)
	}
}
