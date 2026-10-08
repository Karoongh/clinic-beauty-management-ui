package config

import (
	"os"
)

// Config holds only values needed at runtime.
// Secrets never appear in source code.
type Config struct {
	Port        string
	DatabaseURL string
	RedisURL    string
}

// Load reads configuration from environment variables.
func Load() Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return Config{
		Port:        port,
		DatabaseURL: os.Getenv("DATABASE_URL"),
		RedisURL:    os.Getenv("REDIS_URL"),
	}
}
