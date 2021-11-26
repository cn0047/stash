package main

import (
	"fmt"
	"os"

	"github.com/org/repo/app"
	"github.com/org/repo/config"
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

	a := app.New(cfg)
	a.Run()
}
