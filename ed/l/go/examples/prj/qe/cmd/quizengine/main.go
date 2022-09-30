package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"quizengine/app"
	"quizengine/config"
)

var (
	// BuildCommitHash contains hash for GIT commit which was used to build app.
	BuildCommitHash string
)

func main() {
	// Init logger.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Init config.
	cfg, err := config.New()
	if err != nil {
		log.Errorf("failed to create config, err: %+v \n", err)
		os.Exit(1)
	}
	cfg.BuildCommitHash = BuildCommitHash

	// Init application.
	a, err := app.New(cfg)
	if err != nil {
		log.Errorf("failed to create new application, err: %+v \n", err)
		os.Exit(1)
	}

	a.Run()
}
