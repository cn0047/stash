package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signals
		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	log.Fatal(s.ListenAndServe())
}
