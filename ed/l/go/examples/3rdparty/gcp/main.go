package main

import (
	"gcp/storage"
)

const (
	ProjectID  = ""
	SAFilePath = "./sa.json"
)

func main() {
	storage.Run(ProjectID, SAFilePath)
}
