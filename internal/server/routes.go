package server

import (
	"go-wellxs/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func SetupRoutes(app *fiber.App) {
	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false,
		MaxAge:           300,
	}))

	userHandler := handlers.NewUserHandler()
	profileHandler := handlers.NewProfileHandler()

	// API routes
	v1 := app.Group("/api/v1")

	// User routes
	users := v1.Group("/users")
	users.Post("/create-user", userHandler.Create)
	users.Get("/:id", userHandler.GetByID)

	// Profile routes
	profiles := v1.Group("/profiles")
	profiles.Post("/", profileHandler.Create)

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}
