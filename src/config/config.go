package config

import (
	"os"
	"log"
	"github.com/joho/godotenv"
)

// Config - model environment variables
type Config struct {
	Port string
}

// InitEnv - ...
func InitEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

// GetConfig - ...
func GetConfig() *Config {
	cfg := &Config {}

	cfg.Port = getEnv("PORT", ":5000")

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}