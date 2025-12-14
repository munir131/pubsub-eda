package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"pubsub-eda/go/internal/config"
	"pubsub-eda/go/internal/event"
	"pubsub-eda/go/internal/jobs"
	"pubsub-eda/go/internal/platform/pubsub"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Parse CLI Args
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// 3. Setup Context with Cancellation (Graceful Shutdown)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\nReceived shutdown signal")
		cancel()
	}()

	// 4. Initialize Platform (Pub/Sub Client)
	client, err := pubsub.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		log.Fatalf("Failed to create pubsub client: %v", err)
	}
	defer client.Close()

	command := os.Args[1]
	switch command {
	case "publish":
		runPublish(ctx, client, cfg)
	case "listen":
		runListen(ctx, client, cfg)
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func runPublish(ctx context.Context, publisher event.Publisher, cfg *config.Config) {
	fs := flag.NewFlagSet("publish", flag.ExitOnError)
	fs.Parse(os.Args[2:])
	
	args := fs.Args()
	if len(args) < 1 {
		fmt.Println("Usage: app publish <message>")
		os.Exit(1)
	}
	messageBody := args[0]

	msg := &event.Message{
		Data: []byte(messageBody),
	}

	fmt.Printf("Publishing to topic: %s\n", cfg.TopicName)
	if err := publisher.Publish(ctx, cfg.TopicName, msg); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}
	fmt.Printf("Message published! ID: %s\n", msg.ID)
}

func runListen(ctx context.Context, subscriber event.Subscriber, cfg *config.Config) {
	// Initialize Dispatcher
	dispatcher := event.NewDispatcher(subscriber)

	// Register Handlers
	loggerJob := &jobs.LoggerJob{Logger: log.New(os.Stdout, "", log.LstdFlags)}
	dispatcher.Register(cfg.SubscriptionName, loggerJob)

	fmt.Printf("Starting dispatcher... Listening on subscription: %s\n", cfg.SubscriptionName)
	if err := dispatcher.Start(ctx); err != nil {
		// If Start returns error (other than context canceled), log it
		if ctx.Err() == nil {
			log.Fatalf("Dispatcher failed: %v", err)
		}
	}
	fmt.Println("Dispatcher stopped.")
}

func printUsage() {
	fmt.Println("Usage: app [publish|listen] [args]")
}
