package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It is HTTPS!"))
	})
	http.ListenAndServeTLS(":8000", "certificate.pem", "private.key", nil)
}
