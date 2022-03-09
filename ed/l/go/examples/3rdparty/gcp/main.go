package main

import (
	"gcp/bigtable"
	"gcp/iam"
	"gcp/pubsub"
	"gcp/storage"
)

const (
	ProjectID          = ""
	SAFilePath         = "./sa.json"
	BigTableInstanceID = ""
	BucketName         = ""
	PubSubSubscription = "test-1"
)

func main() {
	pubsub.Run(ProjectID, PubSubSubscription)
	return
	bigtable.Run(ProjectID, BigTableInstanceID, SAFilePath)
	iam.Run(ProjectID, SAFilePath)
	storage.Run(ProjectID, SAFilePath, BucketName)
}
