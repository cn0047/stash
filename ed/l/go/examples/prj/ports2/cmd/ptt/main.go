package main

import (
	"os"
	"ptt/app"

	"github.com/sirupsen/logrus"

	"ptt/config"
)

func main() {
	// Init logger.
	log := logrus.New()
	log.Out = os.Stdout
	log.Formatter = &logrus.JSONFormatter{}

	// Init config.
	cfg, err := config.New()
	if err != nil {
		log.Errorf("failed to create config, err: %+v \n", err)
		os.Exit(1)
	}

	// Init application.
	a, err := app.New(cfg, log)
	if err != nil {
		log.Errorf("failed to create new application, err: %+v \n", err)
		os.Exit(1)
	}

	a.Run()
}
