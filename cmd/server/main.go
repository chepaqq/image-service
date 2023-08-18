package main

import (
	"log"
	"net/http"
	"os"

	"github.com/chepaqq99/jungle-task/pkg/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/images").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/login").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/register").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/upload-picture").HandlerFunc(nil)

	cfgDB := db.Config{
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Port:     os.Getenv("POSTGRES_PORT"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Host:     os.Getenv("POSTGRES_HOST"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	}
	_, err := db.ConnectPostgres(cfgDB)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Starting server on :8000")
	err = http.ListenAndServe(":8000", r)
	log.Fatal(err)
}
