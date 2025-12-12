package pubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"pubsub-eda/go/internal/event"
)

// Client wraps the Google Cloud Pub/Sub client.
type Client struct {
	client *pubsub.Client
}

// NewClient creates a new Pub/Sub client.
func NewClient(ctx context.Context, projectID string) (*Client, error) {
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %w", err)
	}
	return &Client{client: client}, nil
}

// Close closes the underlying client.
func (c *Client) Close() error {
	return c.client.Close()
}

// Publish publishes a message to a topic.
func (c *Client) Publish(ctx context.Context, topicID string, msg *event.Message) error {
	t := c.client.Topic(topicID)
	// Ensure the topic exists or handle error (for this exercise we assume it exists
	// or is created by infrastructure/setup).

	result := t.Publish(ctx, &pubsub.Message{
		Data:       msg.Data,
		Attributes: msg.Attributes,
	})

	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("publish.Get: %w", err)
	}

	msg.ID = id
	return nil
}

// Subscribe listens for messages on a subscription and invokes the handler.
func (c *Client) Subscribe(ctx context.Context, subID string, handler event.Handler) error {
	sub := c.client.Subscription(subID)
	
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		eventMsg := &event.Message{
			ID:         msg.ID,
			Data:       msg.Data,
			Attributes: msg.Attributes,
		}

		// Handle the message
		if err := handler.Handle(ctx, eventMsg); err != nil {
			// Nack the message if processing failed
			msg.Nack()
			return
		}

		// Ack the message if successful
		msg.Ack()
	})

	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}

	return nil
}
