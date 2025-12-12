package event

import "context"

// Message represents a generic event message.
type Message struct {
	ID         string
	Data       []byte
	Attributes map[string]string
}

// Ack acknowledges the message processing.
type Ack func()

// Nack explicitly denies the message processing.
type Nack func()

// Publisher defines the interface for publishing messages.
type Publisher interface {
	Publish(ctx context.Context, topic string, msg *Message) error
	Close() error
}

// Handler handles an incoming message.
type Handler interface {
	Handle(ctx context.Context, msg *Message) error
}

// Subscriber defines the interface for subscribing to topics.
type Subscriber interface {
	// Subscribe starts listening for messages on the given subscription.
	// It should block until the context is cancelled or an error occurs.
	Subscribe(ctx context.Context, subscriptionID string, handler Handler) error
	Close() error
}
