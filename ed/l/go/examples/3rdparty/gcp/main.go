package main

import (
	"gcp/iam"
	"gcp/storage"
)

const (
	ProjectID  = ""
	BucketName = ""
	SAFilePath = "./sa.json"
)

func main() {
	iam.Run(ProjectID, SAFilePath)
	return
	storage.Run(ProjectID, SAFilePath, BucketName)
}
