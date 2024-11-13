// internal/interfaces/handlers/user_handler.go
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/treewalkr/gymtrack/internal/application"
	"github.com/treewalkr/gymtrack/internal/domain"
)

type UserHandler struct {
	service *application.UserService
}

func NewUserHandler(service *application.UserService) *UserHandler {
	return &UserHandler{service: service}
}

type RegisterUserInput struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=trainer client"`
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var input RegisterUserInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	user := &domain.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password, // Remember to hash passwords in a real app
		Role:     input.Role,
	}

	err := h.service.RegisterUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.JSON(user)
}
