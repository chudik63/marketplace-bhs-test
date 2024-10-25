package main

import (
	"log"
	"marketplace-bhs-test/internal/delivery/http"
	"marketplace-bhs-test/internal/infrastructure"
	"marketplace-bhs-test/internal/infrastructure/database"
	"marketplace-bhs-test/internal/infrastructure/hash"
	"marketplace-bhs-test/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := infrastructure.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	_ = db

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)
	userService := service.NewUserService(hasher)

	router := gin.Default()
	http.NewUserHandler(router, userService)

	if err := router.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
