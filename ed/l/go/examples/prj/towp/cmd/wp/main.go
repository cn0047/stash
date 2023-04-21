package main

import (
	"github.com/to-com/go-telemetry/opencensusx"
	"github.com/to-com/wp/foundation"
	"github.com/to-com/wp/internal"
)

func main() {
	logger := foundation.NewLogger()

	logger.Infow("startup", "status", "initializing OT/Instrumentation tracing support")
	opencensusx.InitTelemetryWithServiceName(logger, "wp")

	logger.Infow("startup", "status", "initializing application")
	app, err := internal.New(logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err := app.Serve(); err != nil {
		logger.Fatal(err)
	}
}
