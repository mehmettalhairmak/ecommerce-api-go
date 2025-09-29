package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	StripeSecretKey string
	Port            string
	IsProduction    bool
	TLSCertPath     string
	TLSKeyPath      string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Load .env file if exists (for development)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := &Config{
		StripeSecretKey: os.Getenv("STRIPE_SECRET_KEY"),
		Port:            getEnvOrDefault("PORT", "4242"),
		IsProduction:    os.Getenv("PRODUCTION") == "true",
		TLSCertPath:     getEnvOrDefault("TLS_CERT_PATH", "cert.pem"),
		TLSKeyPath:      getEnvOrDefault("TLS_KEY_PATH", "key.pem"),
	}

	// Validate required configuration
	if cfg.StripeSecretKey == "" {
		log.Fatal("You need to set STRIPE_SECRET_KEY environment variable")
	}

	return cfg
}

// getEnvOrDefault returns environment variable value or default value
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
