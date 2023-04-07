package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	port := "8080"
	addr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/x", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("EP: x"))
	})
	mux.HandleFunc("/v1/y", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("EP: y"))
	})
	withLogger := NewLog(mux)
	withTime := NewTime(withLogger)

	server := http.Server{
		Addr:         addr,
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		Handler:      withTime,
	}

	fmt.Printf("start app \n")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("failed to serve server, err: %+v \n", err)
	}

	fmt.Printf("stop app \n")
}

type Log struct {
	h http.Handler
}

func (l *Log) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[INFO] %v %v \n", r.Method, r.URL.Path)
	l.h.ServeHTTP(w, r)
}

func NewLog(h http.Handler) *Log {
	return &Log{h: h}
}

type Time struct {
	h http.Handler
}

func (t *Time) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	before := time.Now()
	t.h.ServeHTTP(w, r)
	after := time.Now()
	took := after.Sub(before)
	fmt.Printf("[took] %v \n", took)
}

func NewTime(h http.Handler) *Time {
	return &Time{h: h}
}
