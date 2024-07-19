package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req [%s]\n", time.Now())
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte("502 Bad Gateway Error.\n"))
	})

	p := ":8080"
	http.ListenAndServe(p, nil)
}
