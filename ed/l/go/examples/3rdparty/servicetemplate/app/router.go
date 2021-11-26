package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRoutes initializes all HTTP routes.
func InitRoutes(a *App, r *mux.Router) {
	mainHandler := &MainHTTPHandler{App: a}

	r.HandleFunc("/health", mainHandler.HealthHandler).Methods("GET")
}

// MainHTTPHandler holds all handlers for all HTTP routes.
type MainHTTPHandler struct {
	App *App
}

func (h *MainHTTPHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	msg := fmt.Sprintf(`{"status":"ok", "version":"%s"}`, h.App.Config.BuildCommitHash)
	_, _ = w.Write([]byte(msg))
}
