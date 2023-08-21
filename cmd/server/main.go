package main

import (
	"github.com/chepaqq/jungle-task/internal/app"
	"github.com/chepaqq/jungle-task/internal/config"
	"github.com/chepaqq/jungle-task/pkg/logger"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		logger.Fatalf("Failed to load config: %v", err)
	}
	app.Run(cfg)
}
