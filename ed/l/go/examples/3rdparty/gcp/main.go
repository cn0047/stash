package main

import (
	"gcp/storage"
)

const (
	ProjectID  = ""
	BucketName = ""
	SAFilePath = "./sa.json"
)

func main() {
	storage.Run(ProjectID, SAFilePath, BucketName)
}
