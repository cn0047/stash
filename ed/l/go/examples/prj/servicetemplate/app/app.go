package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"github.com/cn007b/servicetemplate/config"
)

// App represents main application struct.
type App struct {
	Config *config.Config

	HTTPAddress string
	HTTPServer  *http.Server
}

// New returns new application instance.
func New(cfg *config.Config) (*App, error) {
	a := &App{}

	// Init logger.
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)

	// Init HTTP routes.
	router := mux.NewRouter()
	InitRoutes(a, router)

	// Init HTTP server.
	a.HTTPAddress = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	a.HTTPServer = &http.Server{Addr: a.HTTPAddress, Handler: router}

	a.Config = cfg

	return a, nil
}

// Run starts application.
// For now this method starts only HTTP server,
// but GRPC server may be added here later if needed.
func (a *App) Run() {
	// Start HTTP server.
	log.Infof("starting HTTP server on: %s", a.HTTPAddress)
	go func() {
		err := a.HTTPServer.ListenAndServe()
		if err != nil {
			log.Errorf("failed to ListenAndServe, err: %+v", err)
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

	log.Info("shutting down HTTP server")
	err := a.HTTPServer.Shutdown(ctx)
	if err != nil {
		log.Errorf("failed to perform shutdown, err: %+v", err)
	}
}
