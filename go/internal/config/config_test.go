package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Save original env vars
	origProjectID := os.Getenv("PROJECT_ID")
	origTopicName := os.Getenv("TOPIC_NAME")
	origSubName := os.Getenv("SUBSCRIPTION_NAME")
	defer func() {
		os.Setenv("PROJECT_ID", origProjectID)
		os.Setenv("TOPIC_NAME", origTopicName)
		os.Setenv("SUBSCRIPTION_NAME", origSubName)
	}()

	t.Run("success", func(t *testing.T) {
		os.Setenv("PROJECT_ID", "test-project")
		os.Setenv("TOPIC_NAME", "test-topic")
		os.Setenv("SUBSCRIPTION_NAME", "test-sub")

		cfg, err := Load()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.ProjectID != "test-project" {
			t.Errorf("expected ProjectID 'test-project', got %q", cfg.ProjectID)
		}
		if cfg.TopicName != "test-topic" {
			t.Errorf("expected TopicName 'test-topic', got %q", cfg.TopicName)
		}
		if cfg.SubscriptionName != "test-sub" {
			t.Errorf("expected SubscriptionName 'test-sub', got %q", cfg.SubscriptionName)
		}
	})

	t.Run("missing project id", func(t *testing.T) {
		os.Unsetenv("PROJECT_ID")
		os.Setenv("TOPIC_NAME", "test-topic")
		os.Setenv("SUBSCRIPTION_NAME", "test-sub")

		_, err := Load()
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
