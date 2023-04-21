package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"

	"github.com/to-com/poc-td/app/mfcconfig"
	"github.com/to-com/poc-td/app/storage"
	"github.com/to-com/poc-td/app/toteassignment"
	"github.com/to-com/poc-td/app/tsc"
	"github.com/to-com/poc-td/config"
)

// App represents main application struct.
type App struct {
	Config         *config.Config
	MFCConfig      mfcconfig.MFCConfig
	TSC            tsc.toServiceCatalog
	ToteAssignment toteassignment.ToteAssignment
	log            *zap.SugaredLogger

	HTTPAddress string
	HTTPServer  *http.Server
}

// New returns new application instance.
func New(cfg *config.Config, log *zap.SugaredLogger) (*App, error) {
	a := &App{Config: cfg, log: log}

	// Init GCP Spanner.
	s, err := storage.NewSpannerStorage(cfg.SpannerDatabase)
	if err != nil {
		return nil, fmt.Errorf("failed to create new spanner storage, err: %w", err)
	}

	// Init TSC service.
	a.TSC = tsc.NewService(log)

	// Init MFCConfig service.
	a.MFCConfig = mfcconfig.New(s, a.TSC, log, cfg.MFCConfigAutoRefreshTime)

	// Init ToteAssignment service.
	a.ToteAssignment, err = toteassignment.NewService(s, a.MFCConfig, log, cfg.DBLogEnabled)
	if err != nil {
		return nil, fmt.Errorf("failed to create new stote assignment service, err: %w", err)
	}

	// Init HTTP routes.
	router := mux.NewRouter()
	err = InitRoutes(a, router)
	if err != nil {
		return nil, fmt.Errorf("failed to init routes, err: %w", err)
	}

	// Init HTTP server.
	a.HTTPAddress = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	a.HTTPServer = &http.Server{Addr: a.HTTPAddress, Handler: router}

	return a, nil
}

// Run starts application.
// For now this method starts only HTTP server,
// but GRPC server may be added here later if needed.
func (a *App) Run() {
	// Start HTTP server.
	a.log.With(zap.String("addr", a.HTTPAddress)).Infof("starting HTTP server")
	go func() {
		err := a.HTTPServer.ListenAndServe()
		if err != nil {
			a.log.Errorf("failed to ListenAndServe, err: %+v", err)
		}
	}()

	// Graceful shutdown HTTP server.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C).
	<-c
	a.Shutdown()

	os.Exit(0)
}

// Shutdown stops application.
func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	a.log.Info("shutting down HTTP server")
	err := a.HTTPServer.Shutdown(ctx)
	if err != nil {
		a.log.Errorf("failed to perform shutdown, err: %+v", err)
	}

	a.log.Info("shutting down MFC config")
	a.MFCConfig.Shutdown(ctx)
}
