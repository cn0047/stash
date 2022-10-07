package main

import (
	"gcp/bigtable"
	"gcp/gsm"
	"gcp/iam"
	"gcp/pubsub"
	"gcp/storage"
)

const (
	ProjectID          = "test-project"
	SAFilePath         = "./sa.json"
	BigTableInstanceID = ""
	BucketName         = ""
	PubSubTopic        = "test-topic"
	PubSubSubscription = "test-subscription"
)

func main() {
	pubsub.Run(ProjectID, PubSubTopic, PubSubSubscription)
	return
	gsm.Run(ProjectID, SAFilePath)
	bigtable.Run(ProjectID, BigTableInstanceID, SAFilePath)
	iam.Run(ProjectID, SAFilePath)
	storage.Run(ProjectID, SAFilePath, BucketName)
}
