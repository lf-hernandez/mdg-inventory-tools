package main

import (
	"os"
	"strings"
)

func loadConfig() Config {
	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Port:        os.Getenv("PORT"),
		CORSOrigins: strings.Split(os.Getenv("CORS_ORIGINS"), ","),
	}
}
