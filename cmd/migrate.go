package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/treewalkr/gymtrack/internal/config"
	"github.com/treewalkr/gymtrack/internal/infrastructure/persistence"
)

func MigrateCmdHandler(cmd *cobra.Command, args []string) {
	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DBName)

	_, err := persistence.NewUserRepository(dsn)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully.")
}
