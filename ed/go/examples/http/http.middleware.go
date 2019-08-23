package main

import (
	"fmt"
	"net/http"
	"time"
)

func middlewareT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("⏱ %s\n", time.Now())
		next.ServeHTTP(w, r)
	})
}
func middlewareR(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("✳️ http method: %s\n", r.Method)
		next.ServeHTTP(w, r)
	})
}

func main() {
	finalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	http.Handle("/", middlewareT(middlewareR(finalHandler)))
	http.ListenAndServe(":8080", nil)
}
