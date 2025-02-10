package handlers

import (
	"go-wellxs/internal/database"
	"go-wellxs/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProfileHandler struct {
	DB *gorm.DB
}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{
		DB: database.DB,
	}
}

func (h *ProfileHandler) Create(c *fiber.Ctx) error {
	var profile models.Profile

	if err := c.BodyParser(&profile); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Check if user exists
	var user models.User
	if err := h.DB.First(&user, profile.UserID).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	result := h.DB.Create(&profile)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create profile",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Profile created successfully",
		"profile": profile,
	})
}
