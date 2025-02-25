package main

import (
	"github.com/chepaqq/image-service/internal/app"
	"github.com/chepaqq/image-service/internal/config"
	"github.com/chepaqq/image-service/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	}
	app.Run(cfg)
}
