package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/images").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/login").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/register").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/upload-picture").HandlerFunc(nil)

	log.Print("Starting server on :8000")
	err := http.ListenAndServe(":8000", r)
	log.Fatal(err)
}
