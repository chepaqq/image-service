package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chepaqq/jungle-task/internal/config"
	"github.com/chepaqq/jungle-task/internal/delivery/api/handler"
	"github.com/chepaqq/jungle-task/internal/repository"
	"github.com/chepaqq/jungle-task/internal/service"
	"github.com/chepaqq/jungle-task/pkg/database"
	"github.com/gorilla/mux"
)

// Run initialize and starts application
func Run() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

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

	// Repos
	userRepository := repository.NewUserRepository(postgresClient)

	// Services
	userService := service.NewUserService(userRepository)

	// Handlers
	userHandler := handler.NewUserHandler(userService)

	// Routes
	router := mux.NewRouter()

	router.HandleFunc("/login", userHandler.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/register", userHandler.SignUp).Methods(http.MethodPost)
	router.HandleFunc("/images", nil).Methods(http.MethodGet)
	router.HandleFunc("/upload-picture", nil).Methods(http.MethodPost)

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
