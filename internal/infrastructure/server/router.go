// internal/infrastructure/server/router.go
package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/treewalkr/gymtrack/internal/interfaces/handlers"
)

func SetupRouter(userHandler *handlers.UserHandler) *fiber.App {
	app := fiber.New()

	api := app.Group("/api")

	user := api.Group("/users")
	user.Post("/register", userHandler.Register)
	user.Get("/:id", userHandler.GetUser)

	// Add more routes as needed

	return app
}
