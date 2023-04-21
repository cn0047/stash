package pubsub

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/to-com/wp/config"
	"go.opencensus.io/trace"
	"go.uber.org/zap"
)

type PubSub struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
	client *pubsub.Client
}

func New(cfg *config.Config, logger *zap.SugaredLogger) (*PubSub, error) {
	client, err := pubsub.NewClient(context.Background(), cfg.ProjectID)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		cfg:    cfg,
		logger: logger,
		client: client,
	}, nil
}

func (p *PubSub) PublishMessage(ctx context.Context, topicID string, message any, attrs map[string]string) (*string, error) {
	ctx, span := trace.StartSpan(ctx, "pubsub."+attrs["event_type"])
	defer span.End()

	messageJSON, jsonError := json.Marshal(message)
	if jsonError != nil {
		return nil, jsonError
	}

	topic := p.client.TopicInProject(topicID, p.cfg.ProjectID)
	defer topic.Stop()

	msg := &pubsub.Message{
		Data:       messageJSON,
		Attributes: attrs,
	}
	result := topic.Publish(ctx, msg)

	p.logger.Infof("publishing message '%s' with attributes '%v' to the topic '%s'", msg.Data, msg.Attributes, topicID)
	id, err := result.Get(ctx)
	if err != nil {
		p.logger.Errorf("failed to publish message: '%s', error: '%v'", messageJSON, zap.Error(err))
		return nil, err
	}
	p.logger.Infof("successfully published event to topic: '%s' with id: '%s'", topicID, id)

	return &id, nil
}
