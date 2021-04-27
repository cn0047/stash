package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	// one()
	readBodyTwice()
	return

	go two()
	go two()

	select {}
}

// go run ~/web/kovpak/gh/ed/go/examples/whatever/compare.index.go
// GODEBUG=scheddetail=1,schedtrace=1000 go run ~/web/kovpak/gh/ed/go/examples/http/http.client.get.go
func two() {
	req, err := http.NewRequest("GET", "http://localhost:8080/?busy", nil)
	if err != nil {
		panic(err.Error())
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Resp: %s\n", body)
}

func one() {
	req, err := http.NewRequest("GET", "http://localhost:8080/get", nil)
	if err != nil {
		panic("[1] " + err.Error())
	}

	q := req.URL.Query()
	q.Add("type", "test")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic("[2] " + err.Error())
	}
	defer resp.Body.Close()

	fmt.Printf("Resp: %+v", resp.Body)
}

func he(err error) {
	if err != nil {
		panic(err)
	}
}

func readBodyTwice() {
	req, err := http.NewRequest("GET", "https://api.github.com/users/cn007b", nil)
	he(err)

	client := &http.Client{}
	res, err := client.Do(req)
	he(err)
	defer res.Body.Close()

	var body bytes.Buffer
	bodyReader := io.TeeReader(res.Body, &body)

	d1, err := ioutil.ReadAll(bodyReader)
	he(err)
	d2, err := ioutil.ReadAll(&body)
	he(err)

	fmt.Printf("%s \n======\n", d1)
	fmt.Printf("%s \n======\n", d2)
}
