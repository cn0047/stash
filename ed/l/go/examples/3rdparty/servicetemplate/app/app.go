package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/cn007b/servicetemplate/config"
)

// App represents main application struct.
type App struct {
	Config *config.Config
}

// New returns new application instance.
func New(cfg *config.Config) (*App, error) {
	a := &App{Config: cfg}

	return a, nil
}

// Run starts application.
// For now this method starts only HTTP server,
// but GRPC server may be added here later if needed.
func (a *App) Run() {
	// Init HTTP routes.
	router := mux.NewRouter()
	InitRoutes(a, router)

	// Start HTTP server.
	addr := fmt.Sprintf("%s:%d", a.Config.Host, a.Config.Port)
	srv := &http.Server{Addr: addr, Handler: router}
	log.Printf("starting HTTP server on: %s\n", addr)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Printf("failed to ListenAndServe, err: %+v \n", err)
		}
	}()

	// Graceful shutdown HTTP server.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C).
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	err := srv.Shutdown(ctx)
	if err != nil {
		fmt.Printf("failed to perform shutdown, err: %+v \n", err)
	}

	log.Println("shutting down HTTP server")
	os.Exit(0)
}
