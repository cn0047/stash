package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("req [%s]\n", time.Now())
		w.Write([]byte("Hello world!\n"))
	})
	p := ":4433"
	c := "/Users/k/web/kovpak/gh/ed/l/js.nodejs/nodejs.express/examples/coursera.passport/bin/httpsKeys/certificate.pem"
	k := "/Users/k/web/kovpak/gh/ed/l/js.nodejs/nodejs.express/examples/coursera.passport/bin/httpsKeys/private.key"
	http.ListenAndServeTLS(p, c, k, nil)
}
