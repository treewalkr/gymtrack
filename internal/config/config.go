// internal/config/config.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func LoadConfig() Config {
	env := os.Getenv("ENV")

	// Load environment variables from the appropriate file based on ENV variable
	if env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, loading default environment variables.")
		}
	} else {
		err := godotenv.Load(".env." + env)
		if err != nil {
			log.Printf("No .env.%s file found, loading default environment variables.\n", env)
		}
	}

	viper.AutomaticEnv()

	// Set default values (Optional)
	viper.SetDefault("SERVER_PORT", "5555")
	viper.SetDefault("DATABASE_HOST", "localhost")
	viper.SetDefault("DATABASE_PORT", "5432")
	viper.SetDefault("DATABASE_USER", "gymtrack_user")
	viper.SetDefault("DATABASE_PASSWORD", "securepassword")
	viper.SetDefault("DATABASE_DBNAME", "gymtrack_db")

	// Map environment variables to our Config structure fields
	var cfg Config

	cfg.Server.Port = viper.GetString("SERVER_PORT")
	cfg.Database.Host = viper.GetString("DATABASE_HOST")
	cfg.Database.Port = viper.GetString("DATABASE_PORT")
	cfg.Database.User = viper.GetString("DATABASE_USER")
	cfg.Database.Password = viper.GetString("DATABASE_PASSWORD")
	cfg.Database.DBName = viper.GetString("DATABASE_DBNAME")

	return cfg
}
