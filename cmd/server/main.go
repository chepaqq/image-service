package main

import (
	"log"
	"net/http"

	"github.com/chepaqq99/jungle-task/pkg/db"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/images").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/login").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/register").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/upload-picture").HandlerFunc(nil)

	cfgDB := db.Config{
		Username: "postgres",
		Password: "qwerty",
		Port:     "5432",
		DBName:   "test_db",
		Host:     "postgres",
		SSLMode:  "disable",
	}
	_, err := db.ConnectPostgres(cfgDB)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Starting server on :8000")
	err = http.ListenAndServe(":8000", r)
	log.Fatal(err)
}
