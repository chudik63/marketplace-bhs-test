package main

import (
	"log"
	"marketplace-bhs-test/internal/app"
	"marketplace-bhs-test/internal/infrastructure"
)

func main() {
	cfg, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app.Run(cfg)
}
