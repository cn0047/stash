package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	one()
}

func one() {
	url := "https://itismonitoring.appspot.com/"
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{
		Transport: &http.Transport{
			DialContext:         (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
			TLSHandshakeTimeout: 2 * time.Second,
		},
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body length:", len(string(body)))

	err = res.Body.Close()
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}
}
