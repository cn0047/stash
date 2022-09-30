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

	"quizengine/app/quizengine"
	"quizengine/app/storage"
	"quizengine/config"
)

// App represents main application struct.
type App struct {
	Config     *config.Config
	QuizEngine *quizengine.Service

	HTTPAddress string
	HTTPServer  *http.Server
}

// New returns new application instance.
func New(cfg *config.Config) (*App, error) {
	a := &App{Config: cfg}

	// Init GCP Spanner.
	s, err := storage.NewSpannerStorage(cfg.SpannerDatabase)
	if err != nil {
		return nil, fmt.Errorf("failed to create new spanner storage, err: %w", err)
	}

	// Init QuizEngine service.
	a.QuizEngine, err = quizengine.NewService(s)
	if err != nil {
		return nil, fmt.Errorf("failed to create new quiz engine service, err: %w", err)
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

// Run starts application HTTP server.
func (a *App) Run() {
	// Start HTTP server.
	log.Infof("starting HTTP server")
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
