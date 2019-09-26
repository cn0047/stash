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

	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req [%s]\n", time.Now())
		w.Write([]byte("Hello world!\n"))
	})

	s := http.Server{
		Addr:         p,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
		Handler:      router,
	}
	log.Fatal(s.ListenAndServe())
}
