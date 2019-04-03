package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	p := ":8080"
	if len(os.Args[1]) > 1 {
		p = ":" + os.Args[1]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req [%s]\n", time.Now())
		w.Write([]byte("Hello world!\n"))
	})
	http.ListenAndServe(p, nil)
}
