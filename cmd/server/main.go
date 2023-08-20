package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chepaqq99/jungle-task/internal/config"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/images").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/login").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/register").HandlerFunc(nil)
	r.Methods(http.MethodPost).Path("/upload-picture").HandlerFunc(nil)

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(cfg)
	log.Print("Starting server on :", cfg.Server.Port)
	err = http.ListenAndServe(":"+cfg.Server.Port, r)
	log.Fatal(err)
}
