package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/treewalkr/gymtrack/internal/config"
	"github.com/treewalkr/gymtrack/pkg/logger"
)

func ServeCmdHandler(cmd *cobra.Command, args []string) {
	// Initialize Logger
	logger.InitLogger()
	defer func() {
		if err := logger.Logger.Sync(); err != nil {
			log.Printf("Error syncing logger: %v", err)
		}
	}()

	// Initialize Dependency Injection
	cfg := config.LoadConfig()
	app, err := InitializeServer(
		// Construct DSN
		"host=" + cfg.Database.Host +
			" port=" + cfg.Database.Port +
			" user=" + cfg.Database.User +
			" password=" + cfg.Database.Password +
			" dbname=" + cfg.Database.DBName +
			" sslmode=disable",
	)
	if err != nil {
		logger.Logger.Fatal("Failed to initialize server", logger.Error(err))
	}

	// Start Server
	if err := app.Listen(":" + cfg.Server.Port); err != nil {
		logger.Logger.Fatal("Failed to start server", logger.Error(err))
	}
}
