package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func debug(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprintf(w, "%s", dump); err != nil {
		log.Printf("🟥 error: %#v\n", err)
	}
	log.Printf("✳️\n%s", dump)
}

func main() {
	http.HandleFunc("/", debug)

	http.ListenAndServe(":8080", nil)
}
