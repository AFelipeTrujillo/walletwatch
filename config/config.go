package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBDriver   string
	DBName     string
	ServerPort string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("Warning: .env file not found, using system environment variables")
	}

	config := &Config{
		DBDriver:   getEnv("DB_DRIVER", "sqlite"),
		DBName:     getEnv("DB_NAME", "walletwatch.db"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	return config
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
