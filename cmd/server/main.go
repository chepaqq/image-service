package main

import (
	"log"

	"github.com/chepaqq/jungle-task/internal/app"
	"github.com/chepaqq/jungle-task/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	app.Run(cfg)
}
