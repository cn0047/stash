package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	p := ":8080"
	if len(os.Args) > 1 {
		p = ":" + os.Args[1]
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req [%s]\n", time.Now())
		w.Write([]byte("Hello world!\n"))
	})

	s := http.Server{
		Addr:         p,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Handler:      mux,
	}
	log.Fatal(s.ListenAndServe())
}
