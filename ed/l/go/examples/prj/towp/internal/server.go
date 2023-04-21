package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (a *Application) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", a.cfg.Port),
		Handler:      a.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  a.cfg.ServerReadTimeout,
		WriteTimeout: a.cfg.ServerWriteTimeout,
	}

	a.logger.Infof("running wp on port: %d", a.cfg.Port)

	// receive any errors returned by the graceful Shutdown() function
	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		a.logger.Infof("shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		shutdownError <- srv.Shutdown(ctx)
	}()

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		a.logger.Fatal("an error occurred while gracefully shutting down server")

		return err
	}

	if err := <-shutdownError; err != nil {
		return err
	}

	a.logger.Infof("wp has been shut down successfully: %d", a.cfg.Port)

	return nil
}
