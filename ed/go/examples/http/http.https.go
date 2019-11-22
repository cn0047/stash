package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprintf(w, "%s", dump); err != nil {
			log.Printf("🟥 error: %#v\n", err)
		}
		log.Printf("✳️\n%s", dump)
	})
	c := "/Users/kovpakvolodymyr/web/kovpak/gh/ed/nodejs/nodejs.express/examples/coursera.passport/bin/httpsKeys/certificate.pem"
	k := "/Users/kovpakvolodymyr/web/kovpak/gh/ed/nodejs/nodejs.express/examples/coursera.passport/bin/httpsKeys/private.key"
	http.ListenAndServeTLS(":4433", c, k, nil)
}
