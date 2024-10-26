package app

import (
	"log"
	"marketplace-bhs-test/internal/auth"
	"marketplace-bhs-test/internal/delivery/http"
	"marketplace-bhs-test/internal/delivery/http/middleware"
	"marketplace-bhs-test/internal/infrastructure"
	"marketplace-bhs-test/internal/infrastructure/database"
	"marketplace-bhs-test/internal/repository"
	"marketplace-bhs-test/internal/service"

	"github.com/gin-gonic/gin"
)

func Run(cfg *infrastructure.Config) {
	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	tokenManager, err := auth.NewManager(cfg.Auth.SecretKey)
	if err != nil {
		log.Fatalf("Failed to create token manager: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	assetRepo := repository.NewAssetRepository(db, userRepo)

	userService := service.NewUserService(userRepo, tokenManager, cfg.Auth.AccessTokenTTL, cfg.Auth.RefreshTokenTTL)
	assetService := service.NewAssetService(assetRepo)

	router := gin.Default()
	middleware := middleware.AuthMiddleware(tokenManager)

	http.NewUserHandler(router, userService, middleware)
	http.NewAssetHandler(router, assetService, middleware)

	if err := router.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
