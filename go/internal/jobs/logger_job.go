package jobs

import (
	"context"
	"fmt"
	"log"

	"pubsub-eda/go/internal/event"
)

// LoggerJob is a simple handler that logs the message.
type LoggerJob struct {
	Logger *log.Logger
}

// Handle processes the incoming message.
func (j *LoggerJob) Handle(ctx context.Context, msg *event.Message) error {
	logMsg := fmt.Sprintf("LoggerJob: Received message ID=%s Data=%s Attributes=%v", msg.ID, string(msg.Data), msg.Attributes)
	
	if j.Logger != nil {
		j.Logger.Println(logMsg)
	} else {
		fmt.Println(logMsg)
	}
	return nil
}
