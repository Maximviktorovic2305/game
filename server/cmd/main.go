package main

import (
	"log"
	"server/internal/config"
	"server/internal/database"
	"server/internal/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	db := database.ConnectDB(cfg)
	defer func() {
		if err := database.CloseDB(db); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	// Run migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Seed questions
	if err := database.SeedQuestions(db); err != nil {
		log.Printf("Warning: Failed to seed questions: %v", err)
	}

	// Setup routes and start server
	e := routes.SetupRoutes(db)
	log.Printf("Server starting on port %s", cfg.Port)
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
