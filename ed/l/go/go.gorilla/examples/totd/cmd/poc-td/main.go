package main

import (
	"fmt"
	"os"

	"github.com/to-com/go-log/zapx"

	"github.com/to-com/poc-td/app"
	"github.com/to-com/poc-td/config"
)

var (
	// BuildCommitHash contains hash for GIT commit which was used to build app.
	BuildCommitHash string
)

func main() {
	// Init logger.
	log, err := zapx.New(zapx.Config{ServiceName: "td"})
	if err != nil {
		fmt.Printf("failed to create logger, err: %+v \n", err)
		os.Exit(1)
	}

	// Init config.
	cfg, err := config.New()
	if err != nil {
		log.Errorf("failed to create config, err: %+v \n", err)
		os.Exit(1)
	}
	cfg.BuildCommitHash = BuildCommitHash

	// Init application.
	a, err := app.New(cfg, log)
	if err != nil {
		log.Errorf("failed to create new application, err: %+v \n", err)
		os.Exit(1)
	}

	a.Run()
}
