package config

import (
	"log"
	"os"
	"strings"
)

type Config struct {
	Database    DatabaseConfig
	Port        string
	CORSOrigins []string
	JWTSecret   string
}

type DatabaseConfig struct {
	URL      string
	MaxRetry int
}

func LoadConfig() Config {
	return Config{
		Database: DatabaseConfig{
			URL:      getEnvOrDefault("DATABASE_URL", ""),
			MaxRetry: 5,
		},
		Port:        getEnvOrDefault("PORT", "8000"),
		CORSOrigins: strings.Split(getEnvOrDefault("CORS_ORIGINS", ""), ","),
		JWTSecret:   getEnvOrFatal("JWT_SECRET"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set", key)
	}
	return value
}
