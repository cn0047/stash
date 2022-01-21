package test

import (
	"os"
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/cn007b/servicetemplate/app"
	"github.com/cn007b/servicetemplate/config"
)

var (
	testApp *app.App
)

// InitApp initializes test application.
func InitApp(t *testing.T, cfg *config.Config) {
	testApp = &app.App{Config: cfg}

	// Init logger.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

// GetApp gets test application.
func GetApp(t *testing.T) *app.App {
	if testApp == nil {
		InitApp(t, GetConfig())
	}

	return testApp
}

// GetConfig gets config for test application.
func GetConfig() *config.Config {
	cfg := &config.Config{
		Env: "test",
	}

	return cfg
}
