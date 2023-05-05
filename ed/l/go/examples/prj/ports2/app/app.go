package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"

	"ptt/app/ports"
	"ptt/app/storage"
	"ptt/config"
)

// App represents main application struct.
type App struct {
	Config       *config.Config
	Log          *logrus.Logger
	Validator    *validator.Validate
	Storage      storage.Storage
	PortsService *ports.PortsService

	HTTPAddress string
	HTTPServer  *http.Server
}

// New returns new application instance.
func New(cfg *config.Config, log *logrus.Logger) (*App, error) {
	a := &App{Config: cfg, Log: log}

	// Init default validator.
	a.Validator = validator.New()

	// Init Storage service.
	if cfg.RedisHost != "" && cfg.RedisPort != "" {
		a.Storage = storage.NewRedisStorage(cfg.RedisHost, cfg.RedisPort, cfg.RedisPassword, cfg.RedisDB)
	} else {
		// @TDB: Default storage.
		a.Storage = storage.NewInMemoryStorage()
	}

	// Init Ports service.
	a.PortsService = ports.New(a.Storage)

	// Init HTTP routes.
	router := mux.NewRouter()
	err := InitRoutes(a, router)
	if err != nil {
		return nil, fmt.Errorf("failed to init routes, err: %w", err)
	}

	// Init HTTP server.
	a.HTTPAddress = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	a.HTTPServer = &http.Server{Addr: a.HTTPAddress, Handler: router}

	return a, nil
}

// Run starts application.
func (a *App) Run() {
	// Start HTTP server.
	a.Log.Infof("starting HTTP server on addr: %s", a.HTTPAddress)
	go func() {
		err := a.HTTPServer.ListenAndServe()
		if err != nil {
			a.Log.Errorf("failed to ListenAndServe, err: %+v", err)
		}
	}()

	// Graceful shutdown HTTP server.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-c
	a.Shutdown()

	os.Exit(0)
}

// Shutdown stops application.
func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	a.Log.Info("shutting down HTTP server")
	err := a.HTTPServer.Shutdown(ctx)
	if err != nil {
		a.Log.Errorf("failed to perform shutdown, err: %+v", err)
	}
}
