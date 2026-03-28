package bootstrap

import (
	"gia-starter-app-V1/internal/delivery/http"
	"gia-starter-app-V1/internal/infrastructure/config"
	"gia-starter-app-V1/internal/infrastructure/database"
	"gia-starter-app-V1/internal/infrastructure/logger"

	"github.com/gin-gonic/gin"
)

func InitApp() {
	// Initialize Logger
	logger.InitLogger()
	logger.Info("Starting GIA Starter App...")

	// Load Configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		logger.Fatal("Failed to load configuration")
	}

	// Initialize Database
	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		logger.Fatal("Database connection failed")
	}

	router := gin.Default()

	// Setup Main Router (aggregates module routes) + pass db untuk module wiring
	http.SetupRouter(router, db)

	router.Run(":8081")
}
