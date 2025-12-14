package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration.
type Config struct {
	ProjectID        string
	TopicName        string
	SubscriptionName string
}

// Load loads the configuration from environment variables.
// It optionally loads from a .env file if present.
func Load() (*Config, error) {
	// Attempt to load .env file, but don't fail if it doesn't exist
	_ = godotenv.Load()

	cfg := &Config{
		ProjectID:        os.Getenv("PROJECT_ID"),
		TopicName:        os.Getenv("TOPIC_NAME"),
		SubscriptionName: os.Getenv("SUBSCRIPTION_NAME"),
	}

	if cfg.ProjectID == "" {
		return nil, fmt.Errorf("PROJECT_ID is required")
	}
	if cfg.TopicName == "" {
		return nil, fmt.Errorf("TOPIC_NAME is required")
	}
	if cfg.SubscriptionName == "" {
		return nil, fmt.Errorf("SUBSCRIPTION_NAME is required")
	}

	return cfg, nil
}
