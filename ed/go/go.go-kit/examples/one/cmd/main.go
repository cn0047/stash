package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	appEnv := "local"
	cfg, err := service.InitConfig(appEnv)
	if err != nil {
		log.Fatal("failed to init config: " + err.Error())
	}

	srv := service.NewService()
	srv = service.NewLoggingMiddleware(srv)
	srv = service.NewInstrumentingMiddleware(srv)
	handler := service.NewHTTPHandler(srv)

	if h, ok := handler.(*chi.Mux); ok {
		h.Get("/about", func(w http.ResponseWriter, r *http.Request) {
			json.NewEncoder(w).Encode(map[string]string{
				"code": "200",
			})
		})
		h.Get("/health", func(w http.ResponseWriter, r *http.Request) {})
	}

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
