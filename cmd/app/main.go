package main

import (
	"log"

	"github.com/dorozhkoanton/battletoads/config"
	"github.com/dorozhkoanton/battletoads/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
