package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

type webRequest struct {
	r      *http.Request
	w      http.ResponseWriter
	doneCh chan struct{}
}

var (
	requestCh = make(chan *webRequest)
)

var (
	transport = http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
)

func init() {
	http.DefaultClient = &http.Client{Transport: &transport}
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		doneCh := make(chan struct{})
		requestCh <- &webRequest{r: r, w: w, doneCh: doneCh}
		<-doneCh
	})

	go processRequests()

	go http.ListenAndServeTLS(":2000", "cert.pem", "key.pem", nil)

	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()
}
