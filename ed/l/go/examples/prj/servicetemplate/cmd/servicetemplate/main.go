package main

import (
	"fmt"
	"os"

	"github.com/cn007b/servicetemplate/app"
	"github.com/cn007b/servicetemplate/config"
)

var (
	// BuildCommitHash contains hash for GIT commit which was used to build app.
	BuildCommitHash string
)

func main() {
	cfg, err := config.New()
	if err != nil {
		fmt.Printf("failed to create config, err: %+v \n", err)
		os.Exit(1)
	}
	cfg.BuildCommitHash = BuildCommitHash

	a, err := app.New(cfg)
	if err != nil {
		fmt.Printf("failed to create new application, err: %+v \n", err)
		os.Exit(1)
	}

	a.Run()
}
