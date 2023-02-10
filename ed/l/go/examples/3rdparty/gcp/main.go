package main

import (
	"gcp/bigtable"
	"gcp/gsm"
	"gcp/iam"
	"gcp/iap"
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
	iap.Run(ProjectID)
	return
	pubsub.Run(ProjectID, PubSubTopic, PubSubSubscription)
	gsm.Run(ProjectID, SAFilePath)
	bigtable.Run(ProjectID, BigTableInstanceID, SAFilePath)
	iam.Run(ProjectID, SAFilePath)
	storage.Run(ProjectID, SAFilePath, BucketName)
}
