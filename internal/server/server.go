package server

import (
	"github.com/gofiber/fiber/v2"

	"go-wellxs/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go-wellxs",
			AppName:      "go-wellxs",
		}),

		db: database.New(),
	}

	return server
}
