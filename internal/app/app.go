package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/chepaqq/jungle-task/internal/config"
	authHandler "github.com/chepaqq/jungle-task/internal/delivery/api/handler/auth"
	authRepository "github.com/chepaqq/jungle-task/internal/repository/auth"
	authService "github.com/chepaqq/jungle-task/internal/service/auth"
	"github.com/chepaqq/jungle-task/pkg/database/postgres"
	"github.com/gorilla/mux"
)

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
	postgresClient, err := postgres.ConnectPostgres(postgresURL)
	if err != nil {
		log.Fatal(err)
	}

	// Repos
	authRepository := authRepository.NewRepository(postgresClient)

	// Services
	authService := authService.NewService(authRepository)

	// Handlers
	authHandler := authHandler.NewHandler(*authService)

	// Routes
	router := mux.NewRouter()

	router.HandleFunc("/login", authHandler.SignIn).Methods(http.MethodPost)
	router.HandleFunc("/register", authHandler.SignUp).Methods(http.MethodPost)
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
