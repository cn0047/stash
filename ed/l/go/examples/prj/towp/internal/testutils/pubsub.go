package testutils

import (
	"context"
	"sync"
	"time"

	"cloud.google.com/go/pubsub"
)

var GetPubSubClient = func(ctx context.Context, projectId string) (*pubsub.Client, error) {
	return pubsub.NewClient(ctx, projectId)
}

func WaitForPubSubMessage(
	projectID string,
	subscriptionID string,
	matcher func(*pubsub.Message) bool,
) (*pubsub.Message, error) {
	ctx := context.Background()
	client, err := GetPubSubClient(ctx, projectID)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	subscriptionCtx, cancelSubscriptionCtx := context.WithTimeout(ctx, 5*time.Second)
	subscription := client.SubscriptionInProject(subscriptionID, projectID)
	var foundMessage *pubsub.Message
	var lock sync.Mutex
	err = subscription.Receive(subscriptionCtx,
		func(_ context.Context, message *pubsub.Message) {
			if matcher(message) {
				message.Ack()
				lock.Lock()
				foundMessage = message
				lock.Unlock()
				cancelSubscriptionCtx()
				return
			}

			message.Nack()
		},
	)

	return foundMessage, err
}
