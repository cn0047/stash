package app

import (
	"github.com/gorilla/mux"
)

// InitRoutes initializes all HTTP routes.
func InitRoutes(a *App, rootRouter *mux.Router) error {
	mainHandler := &MainHTTPHandler{App: a}

	// Init root routes.
	rootRouter.HandleFunc("/", mainHandler.HealthHandler).Methods("GET")
	rootRouter.HandleFunc("/health", mainHandler.HealthHandler).Methods("GET")

	// Init ports API routes.
	r := rootRouter.PathPrefix("/v1").Subrouter()
	r.HandleFunc("/ports", mainHandler.CreatePortHandler).Methods("POST")
	r.HandleFunc("/ports/{portId}", mainHandler.UpdatePortHandler).Methods("PUT")

	rootRouter.Handle("/", rootRouter)

	return nil
}
