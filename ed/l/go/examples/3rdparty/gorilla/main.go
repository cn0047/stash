package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
}

func main() {
	a := &App{}
	router := mux.NewRouter()
	InitRoutes(a, router)

	addr := fmt.Sprintf("%s:%s", "localhost", "9000")
	server := &http.Server{Addr: addr, Handler: router}
	log.Printf("starting HTTP server  %+v", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("failed to ListenAndServe, err: %+v", err)
	}
}

func InitRoutes(a *App, rootRouter *mux.Router) {
	mainHandler := &MainHTTPHandler{App: a}

	rootRouter.HandleFunc("/health", mainHandler.HealthHandler).Methods("GET")
	// // Swagger UI.
	// fs := http.StripPrefix("/apidocs/", http.FileServer(http.Dir("./assets/apidocs/")))
	// rootRouter.PathPrefix("/apidocs/").Handler(fs)

	// Init API routes with middlewares.
	r := rootRouter.PathPrefix("/v1").Subrouter()
	r.Use(authMiddleware)
	r.HandleFunc("/echo", mainHandler.GetEchoHandler).Methods("GET")

	rootRouter.Handle("/", rootRouter)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Token")
		if token == "" {
			fmt.Printf("[authMiddleware] missing X-Token \n")
		}

		next.ServeHTTP(w, r)
	})
}

type MainHTTPHandler struct {
	App *App
}

func (h *MainHTTPHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	h.success200(w, map[string]string{"status": "ok", "version": "0"})
}

func (h *MainHTTPHandler) GetEchoHandler(w http.ResponseWriter, r *http.Request) {
	h.success200(w, map[string]string{"echo": "true", "uri": r.RequestURI})
}

func (h *MainHTTPHandler) success200(w http.ResponseWriter, data interface{}) {
	res, err := json.Marshal(data)
	if err != nil {
		h.error500(w, "Failed to marshal response data.", err)
		return
	}

	_, _ = w.Write(res)
}

func (h *MainHTTPHandler) error400(w http.ResponseWriter, msg string, ref string) {
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"%s","type":"bad request"}]}`, msg, ref)
	_, _ = w.Write([]byte(res))
}

func (h *MainHTTPHandler) error500(w http.ResponseWriter, msg string, err error) {
	log.Printf("%s, err: %v", msg, err)
	res := fmt.Sprintf(`{"errors":[{"error":"%s","ref":"","type":"internal error"}]}`, msg)
	_, _ = w.Write([]byte(res))
}
