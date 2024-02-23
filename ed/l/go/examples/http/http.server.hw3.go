package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET 127.0.0.1/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello 1")
	})
	mux.HandleFunc("GET localhost/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello 2")
	})

	fmt.Println("Open: http://localhost:8080")

	server := &http.Server{Addr: ":8080", Handler: mux}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
