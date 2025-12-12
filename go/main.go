package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

// TODO: Replace with your Google Cloud project ID.
const projectID = "pubsub-eda"
const topicName = "test"
const subscriptionName = "test-sub"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [publish|listen] [message]")
		return
	}

	command := os.Args[1]

	switch command {
	case "publish":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go publish <message>")
			return
		}
		message := os.Args[2]
		if err := publishMessage(projectID, topicName, message); err != nil {
			log.Fatalf("Failed to publish message: %v", err)
		}
	case "listen":
		if err := listenForMessages(projectID, subscriptionName); err != nil {
			log.Fatalf("Failed to listen for messages: %v", err)
		}
	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Usage: go run main.go [publish|listen] [message]")
	}
}

// publishMessage publishes a message to a Pub/Sub topic.
func publishMessage(projectID, topicID, msg string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(msg),
	})

	// Block until the result is returned and a server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		return fmt.Errorf("get: %v", err)
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

// listenForMessages listens for messages on a Pub/Sub subscription.
func listenForMessages(projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)
	fmt.Printf("Listening for messages on %q\n", subID)

	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Printf("Got message: %q\n", string(msg.Data))
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("receive: %v", err)
	}

	return nil
}
