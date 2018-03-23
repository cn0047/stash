package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World from Go!"))
    })

    http.HandleFunc("/health-check", HealthCheckHandler)

    http.ListenAndServe(":8000", nil)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"alive": true}`))
}
