package handlers

import (
	"go-wellxs/internal/database"
	"go-wellxs/internal/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		DB: database.DB,
	}
}

// create user
func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// // Add validation for phone
	// if user.Phone == "" {
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": "Phone number is required",
	// 	})
	// }

	// Check if email already exists
	var existingUser models.User
	if err := h.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email already exists",
		})
	}

	result := h.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    user,
	})
}

// get user detail by ID
func (h *UserHandler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User

	if err := h.DB.First(&user, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}
