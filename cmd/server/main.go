package main

import (
	"log"
	"net/http"
	"time"

	"github.com/chepaqq/jungle-task/internal/config"
	"github.com/chepaqq/jungle-task/internal/delivery/api"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      api.InitRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Print("Starting server on :", cfg.Server.Port)
	err = srv.ListenAndServe()
	log.Fatal(err)
}
