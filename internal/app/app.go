package app

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/chepaqq/jungle-task/internal/config"
	"github.com/chepaqq/jungle-task/internal/delivery/api"
	"github.com/chepaqq/jungle-task/internal/delivery/api/handler"
	"github.com/chepaqq/jungle-task/internal/repository"
	"github.com/chepaqq/jungle-task/internal/service"
	"github.com/chepaqq/jungle-task/pkg/database"
	"github.com/chepaqq/jungle-task/pkg/logger"
	"github.com/chepaqq/jungle-task/pkg/server"
	"github.com/chepaqq/jungle-task/pkg/storage"
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
		logger.Fatalf("Failed to connect to Postgres: %v", err)
	}
	logger.Info("Starting Postgres")
	useSSL, err := strconv.ParseBool(cfg.Minio.SSL)
	if err != nil {
		logger.Fatal(err)
	}

	minioClient, err := storage.ConnectMinio(useSSL, cfg.Minio.Endpoint, cfg.Minio.BucketName, cfg.Minio.BucketLocation)
	if err != nil {
		logger.Fatalf("Failed to connect to Minio: %v", err)
	}

	logger.Info("Starting Minio")
	// Repos
	userRepository := repository.NewUserRepository(postgresClient)
	imageRepository := repository.NewImageRepository(postgresClient, minioClient)

	// Services
	userService := service.NewUserService(userRepository)
	imageService := service.NewImageService(imageRepository)

	// Handlers
	userHandler := handler.NewUserHandler(userService)
	imageHandler := handler.NewImageHandler(imageService)

	// HTTP
	router := api.NewRouter(*userHandler, *imageHandler)
	server := server.New(router, cfg.Server.Port)

	// Waiting signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-interrupt:
		logger.Errorf("Signal interrupt error: " + s.String())
	case err = <-server.Notify():
		logger.Infof("Server notify %v:", err)
	}

	// Shutdown server
	err = server.Shutdown()
	if err != nil {
		logger.Infof("Server shutdown: %v", err)
	}
}
