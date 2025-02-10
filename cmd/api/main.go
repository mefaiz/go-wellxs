package main

import (
	"log"
	"os"

	"go-wellxs/internal/database"
	"go-wellxs/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// Run migrations
	if err := database.MigrateDB(); err != nil {
		log.Fatal(err)
	}

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	server.SetupRoutes(app)

	// Get port from env
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server is running on http://localhost:%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
