// main.go
package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/treewalkr/gymtrack/cmd"
)

func main() {
	rootCmd := &cobra.Command{Use: "gymtrack"}

	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "serve",
			Short: "Start serving the API server",
			Run:   cmd.ServeCmdHandler,
		},
		&cobra.Command{
			Use:   "migrate",
			Short: "Sync schema with the database",
			Run:   cmd.MigrateCmdHandler,
		},
		&cobra.Command{
			Use:   "seed",
			Short: "Create seed data",
			Run:   cmd.SeedCmdHandler,
		},
	)

	// Execute root command
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
