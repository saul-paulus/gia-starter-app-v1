package main

import (
	"gia-starter-app-V1/internal/infrastructure/config"
	"gia-starter-app-V1/internal/infrastructure/database"
	"gia-starter-app-V1/internal/infrastructure/logger"
	"gia-starter-app-V1/internal/seeder"

	"log"
)

func main() {
	// Initialize Logger
	logger.InitLogger()

	// Load Configuration
	cfg, err := config.LoadConfig("configs/config.yaml")

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Database
	db, err := database.InitDB(&cfg.Database)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	// Run Seeders
	log.Println("Starting seeding...")

	if err := seeder.SeedUser(db); err != nil {
		log.Fatalf("User seeding failed: %v", err)
	}

	log.Println("✅ Seeding completed successfully")
}


