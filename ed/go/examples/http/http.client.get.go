package main

import (
	"fmt"
	"net/http"
)

func main() {
	one()
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
