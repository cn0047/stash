package pubsub

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

func Run(projectID, topic, subscription string) {
	ctx := context.Background()
	c, err := getClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	publish(ctx, c, topic)

	err = receive(ctx, c, subscription)
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

func publish(ctx context.Context, c *pubsub.Client, topic string) {
	t := c.Topic(topic)
	result := t.Publish(ctx, &pubsub.Message{
		Data: []byte("Message 1"),
	})
	fmt.Printf("Got result: %+v \n", result)
}

func receive(ctx context.Context, c *pubsub.Client, subscription string) error {
	s := c.Subscription(subscription)
	err := s.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Printf("Got message: %s \n", m.Data)
		m.Ack() // Acknowledge consumed message.
	})
	if err != nil {
		return fmt.Errorf("failed to receive from subscription, err: %w", err)
	}

	return nil
}
