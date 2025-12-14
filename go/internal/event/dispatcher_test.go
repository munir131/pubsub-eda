package event

import (
	"context"
	"testing"
	"time"
)

// MockSubscriber is a mock implementation of Subscriber.
type MockSubscriber struct {
	SubscribeFunc func(ctx context.Context, subscriptionID string, handler Handler) error
	CloseFunc     func() error
}

func (m *MockSubscriber) Subscribe(ctx context.Context, subscriptionID string, handler Handler) error {
	if m.SubscribeFunc != nil {
		return m.SubscribeFunc(ctx, subscriptionID, handler)
	}
	return nil
}

func (m *MockSubscriber) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
	return nil
}

// MockHandler is a mock implementation of Handler.
type MockHandler struct {
	HandleFunc func(ctx context.Context, msg *Message) error
}

func (m *MockHandler) Handle(ctx context.Context, msg *Message) error {
	if m.HandleFunc != nil {
		return m.HandleFunc(ctx, msg)
	}
	return nil
}

func TestDispatcher_RegisterAndStart(t *testing.T) {
	mockSub := &MockSubscriber{
		SubscribeFunc: func(ctx context.Context, subscriptionID string, handler Handler) error {
			if subscriptionID != "test-sub" {
				t.Errorf("expected subscriptionID 'test-sub', got %q", subscriptionID)
			}
			return nil
		},
	}

	d := NewDispatcher(mockSub)
	d.Register("test-sub", &MockHandler{})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	if err := d.Start(ctx); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
