package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/images", nil).Methods(http.MethodGet)
	r.HandleFunc("/login", nil).Methods(http.MethodPost)
	r.HandleFunc("/register", nil).Methods(http.MethodPost)
	r.HandleFunc("/upload-picture", nil).Methods(http.MethodPost)
	return r
}
