package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Vendor string `json:"vendor"`
	Name string `json:"name"`
}

func main() {
	p := Payload{Vendor: "BMW", Name: "X6"}
	jsonPayload, err := json.Marshal(p)
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}

	url := "http://localhost:8080/cars"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}
	defer res.Body.Close()

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))
}
