package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	JWTSecret   string
	Port        string
	Environment string
	FrontendURL string
}

func Load() *Config {
	// Load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/rhoam_together"),
		JWTSecret:   getEnv("JWT_SECRET", "change-this-secret-in-production"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "dev"),
		FrontendURL: getEnv("FRONTEND_URL", "http://localhost:3000"),
	}

	// Warn if using default secrets in production
	if config.Environment == "production" && config.JWTSecret == "change-this-secret-in-production" {
		log.Println("WARNING: Using default JWT secret in production!")
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
