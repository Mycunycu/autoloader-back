package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"
)

// Config - model environment variables
type Config struct {
	Port       string
	DbName     string
	MongoDbURL string
}

// InitEnv - ...
func InitEnv() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("No .env file found")
	}
}

// GetConfig - ...
func GetConfig() *Config {
	cfg := &Config{}

	cfg.Port = getEnv("PORT", ":5000")
	cfg.DbName = getEnv("DBNAME", "autoloader")
	cfg.MongoDbURL = getEnv("MONGODB_URL", "")

	return cfg
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
