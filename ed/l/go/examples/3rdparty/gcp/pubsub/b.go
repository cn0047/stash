package pubsub

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func Run(projectID string, pubSubSubscription string) {
	ctx := context.Background()
	c, err := getClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	err = getFromSubscription(ctx, c, pubSubSubscription)
	if err != nil {
		log.Fatal(err)
	}
}

func getClient(ctx context.Context, projectID string) (*pubsub.Client, error) {
	c, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to create new client, err: %w", err)
	}

	return c, nil
}

func getFromSubscription(ctx context.Context, c *pubsub.Client, pubSubSubscription string) error {
	sub := c.Subscription(pubSubSubscription)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Printf("Got message: %s \n", m.Data)
		m.Ack() // Acknowledge consumed message.
	})
	if err != nil {
		return fmt.Errorf("failed to receive from subscription, err: %w", err)
	}

	return nil
}
