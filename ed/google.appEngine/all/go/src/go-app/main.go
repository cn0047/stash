package main

import (
	"fmt"
	"github.com/thepkg/strings"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	msg := "ago - ok. " + strings.ToUpperFirst("upgraded.")
	w.Write([]byte(msg))
}
