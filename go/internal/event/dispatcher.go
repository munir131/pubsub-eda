package event

import (
	"context"
	"fmt"
	"sync"
)

// Dispatcher manages subscriptions and dispatches messages to handlers.
type Dispatcher struct {
	subscriber Subscriber
	handlers   map[string]Handler
	mu         sync.RWMutex
}

// NewDispatcher creates a new Dispatcher.
func NewDispatcher(subscriber Subscriber) *Dispatcher {
	return &Dispatcher{
		subscriber: subscriber,
		handlers:   make(map[string]Handler),
	}
}

// Register registers a handler for a specific subscription.
func (d *Dispatcher) Register(subscriptionID string, handler Handler) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.handlers[subscriptionID] = handler
}

// Start starts listening on all registered subscriptions.
// It blocks until the context is cancelled.
func (d *Dispatcher) Start(ctx context.Context) error {
	d.mu.RLock()
	subs := make(map[string]Handler, len(d.handlers))
	for k, v := range d.handlers {
		subs[k] = v
	}
	d.mu.RUnlock()

	var wg sync.WaitGroup
	errChan := make(chan error, len(subs))

	for subID, handler := range subs {
		wg.Add(1)
		go func(sid string, h Handler) {
			defer wg.Done()
			fmt.Printf("Starting listener for subscription: %s\n", sid)
			if err := d.subscriber.Subscribe(ctx, sid, h); err != nil {
				errChan <- fmt.Errorf("subscription %s failed: %w", sid, err)
			}
		}(subID, handler)
	}

	// Wait for all subscriptions to finish (which happens on ctx cancel or error)
	// Actually Subscribe blocks, so we wait for them to return.
	// If context is cancelled, they should all return.
	
	// Issue: If one fails immediately, we might want to shut down.
	// For simplicity, we just wait for the context to trigger shutdown in main.
	
	wg.Wait()
	close(errChan)

	// Collect errors
	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("dispatcher stopped with %d errors: %v", len(errs), errs[0])
	}
	return nil
}
