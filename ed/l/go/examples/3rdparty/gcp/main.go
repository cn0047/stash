package main

import (
	"gcp/bigtable"
	"gcp/iam"
	"gcp/storage"
)

const (
	ProjectID          = ""
	BigTableInstanceID = ""
	BucketName         = ""
	SAFilePath         = "./sa.json"
)

func main() {
	bigtable.Run(ProjectID, BigTableInstanceID, SAFilePath)
	return
	iam.Run(ProjectID, SAFilePath)
	storage.Run(ProjectID, SAFilePath, BucketName)
}
