// wire.go
//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/treewalkr/gymtrack/internal/application"
	"github.com/treewalkr/gymtrack/internal/infrastructure/persistence"
	"github.com/treewalkr/gymtrack/internal/infrastructure/server"
	"github.com/treewalkr/gymtrack/internal/interfaces/handlers"
)

func InitializeServer(dsn string) (*fiber.App, error) {
	wire.Build(
		persistence.NewUserRepository, // This returns domain.UserRepository
		application.NewUserService,    // This requires domain.UserRepository
		handlers.NewUserHandler,
		server.SetupRouter,
	)
	return &fiber.App{}, nil
}
