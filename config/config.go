package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort      string
	GinMode         string
	LogLevel        string
	APIKeyTableName string
	APIKeyRegion    string
}

var ApplicationConfig Config

func LoadConfig() *Config {
	_ = godotenv.Load()
	cfg := &Config{
		ServerPort: getEnv("SERVER_PORT", "8000"),
		GinMode:    getEnv("GinMode", "release"),
		LogLevel:   getEnv("LogLevel", "info"),

		APIKeyTableName: getEnv("API_KEY_TABLE", "api_key_store"),
		APIKeyRegion:    getEnv("API_KEY_REGION", "us-east-2"),
	}
	AppConfig := cfg
	return AppConfig

}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
