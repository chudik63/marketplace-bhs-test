package main

import (
	"log"
	"marketplace-bhs-test/internal/app"
	"marketplace-bhs-test/internal/infrastructure"
)

// @title Marketplace-BHS-test
// @version 1.0
// @description test tast for Marketplace-Hive project

// @host localhost:8080

func main() {
	cfg, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	app.Run(cfg)
}
