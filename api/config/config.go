package config

import (
	"os"
	"strings"
)

func LoadConfig() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		CORSOrigins: strings.Split(os.Getenv("CORS_ORIGINS"), ","),
	}
}
