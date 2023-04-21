package pubsub_test

import (
	ps "cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/to-com/wp/config"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal/dto"
	"github.com/to-com/wp/internal/pubsub"
	"github.com/to-com/wp/internal/testutils"
	"math"
	"testing"
	"time"
)

var projectID = "test-project"
var pubsubTopic = "wave-plan.event"
var subscriptionID = "wave-plan.sub"

func preparePubSubClient(t *testing.T) *pubsub.PubSub {
	t.Helper()
	t.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8681")

	cfg, err := config.Load()
	if err != nil {
		t.Fatalf("unable to load config for tests, error: %v", err)
	}
	logger := foundation.NewLogger()
	pubSubClient, err := pubsub.New(cfg, logger)
	if err != nil {
		t.Fatalf("unable to initialize pubsub, error: %v", err)
	}

	return pubSubClient
}

func TestPublishMessage(t *testing.T) {
	pubsubClient := preparePubSubClient(t)
	ctx := testutils.BuildCtx()

	var wp dto.wpResponse
	err := json.Unmarshal([]byte(testutils.ReadFileAsString(t, "./../testdata/wpWithSchedules.json")), &wp)
	if err != nil {
		t.Skipf("Was not able to parse JSON")
	}
	message := dto.wpCreatedEvent{
		CreatedAt: time.Now(),
		wp:  wp,
	}
	pubsubAttributes := map[string]string{
		"env_type":    "dev",
		"event_type":  "wp.Created",
		"retailer_id": "fake-retailer",
		"mfc_id":      "fake-mfc",
		"source":      "wp",
	}

	if err != nil {
		t.Fatalf("Was not able to get pubSubClient: %s", err)
	}

	messageID, err := pubsubClient.PublishMessage(ctx, pubsubTopic, message, pubsubAttributes)
	if err != nil {
		t.Errorf("Error occurred while publishing message '%#v' to the topic: %s", message, err)
		return
	}
	assert.Nil(t, err)
	assert.NotNil(t, messageID)

	receivedMessage, err := testutils.WaitForPubSubMessage(
		projectID,
		subscriptionID,
		func(message *ps.Message) bool {
			return message.ID == *messageID
		},
	)
	assert.Nil(t, err)
	assert.NotNil(t, receivedMessage)
	assert.Equal(t, pubsubAttributes, receivedMessage.Attributes)

	messageJSON, err := json.Marshal(message)
	if err != nil {
		t.Errorf("Error occurs while converting received message to JSON: %s", err)
	}
	assert.Nil(t, err)

	assert.Equal(t, string(messageJSON), string(receivedMessage.Data))
}

func TestPublishMessageBadJson(t *testing.T) {
	pubsubClient := preparePubSubClient(t)

	var wrongMsg = math.Inf(1)
	pubsubAttributes := map[string]string{
		"env_type":    "dev",
		"event_type":  "wp.Created",
		"retailer_id": "fake-retailer",
		"mfc_id":      "fake-mfc",
		"source":      "wp",
	}
	_, err := pubsubClient.PublishMessage(context.Background(), pubsubTopic, wrongMsg, pubsubAttributes)

	assert.Equal(t, "json: unsupported value: +Inf", err.Error())
}
