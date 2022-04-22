package main

import (
	"gcp/bigtable"
	"gcp/gsm"
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
	gsm.Run(ProjectID, SAFilePath)
	return
	pubsub.Run(ProjectID, PubSubSubscription)
	bigtable.Run(ProjectID, BigTableInstanceID, SAFilePath)
	iam.Run(ProjectID, SAFilePath)
	storage.Run(ProjectID, SAFilePath, BucketName)
}
