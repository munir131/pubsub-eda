package jobs

import (
	"bytes"
	"context"
	"log"
	"strings"
	"testing"

	"pubsub-eda/go/internal/event"
)

func TestLoggerJob_Handle(t *testing.T) {
	var buf bytes.Buffer
	logger := log.New(&buf, "", 0)

	job := &LoggerJob{Logger: logger}

	msg := &event.Message{
		ID:   "123",
		Data: []byte("test-data"),
		Attributes: map[string]string{
			"key": "value",
		},
	}

	if err := job.Handle(context.Background(), msg); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	output := buf.String()
	if !strings.Contains(output, "LoggerJob: Received message ID=123 Data=test-data") {
		t.Errorf("expected log output to contain message details, got: %s", output)
	}
}
