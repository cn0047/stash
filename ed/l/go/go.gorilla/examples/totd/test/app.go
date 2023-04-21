package test

import (
	"testing"

	"github.com/to-com/go-log/zapx"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/to-com/poc-td/app"
	"github.com/to-com/poc-td/app/mfcconfig"
	"github.com/to-com/poc-td/app/storage"
	"github.com/to-com/poc-td/app/toteassignment"
	"github.com/to-com/poc-td/config"
)

var (
	testApp *app.App
)

// InitApp initializes test application.
func InitApp(t *testing.T, cfg *config.Config) {
	testApp = &app.App{Config: cfg}

	// Init logger.
	log, err := zapx.New(zapx.Config{ServiceName: "test-td"})
	So(err, ShouldBeNil)

	// Init GCP Spanner.
	s, err := storage.NewSpannerStorage(cfg.SpannerDatabase)
	So(err, ShouldBeNil)

	// Init MFCConfig service.
	testApp.MFCConfig = mfcconfig.New(s, nil, log, cfg.MFCConfigAutoRefreshTime)

	// Init ToteAssignment service.
	testApp.ToteAssignment, err = toteassignment.NewService(s, testApp.MFCConfig, log, cfg.DBLogEnabled)
	So(err, ShouldBeNil)
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
		Env:                      "test",
		SpannerDatabase:          "projects/test-project/instances/test-instance/databases/test-db",
		DBLogEnabled:             false,
		MFCConfigAutoRefreshTime: 0,
	}

	return cfg
}
