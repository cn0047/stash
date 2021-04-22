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
	Name   string `json:"name"`
}

func main() {
	// one()
	two()
}

func one() {
	p := Payload{Vendor: "BMW", Name: "X6"}
	jsonPayload, err := json.Marshal(p)
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}

	url := "http://localhost:8080/cars"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic("RUNTIME-ERROR: " + err.Error())
	}

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))

	err = res.Body.Close()
	if err != nil {
		return panic(fmt.Errorf("failed to close response body, error: %w", err))
	}
}

func two() {
	data := make(map[string]string)
	data["code"] = "200"
	data["status"] = "OK"
	jsonPayload, _ := json.Marshal(data)
	fmt.Printf("%+#\n", jsonPayload)

	// url := "https://realtimelog.herokuapp.com:443/sddjklskj"
	// req, x := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	// req.Header.Set("Content-Type", "application/json")
	// fmt.Printf("%+#\n", x)
}

func formUrlEncoded() {
	parm := url.Values{}
	parm.Add("grant_type", "refresh_token")
	parm.Add("refresh_token", "t")

	u := "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer"
	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(parm.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Basic ht")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("response Status:", res.Status)
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println("response Body:", string(body))

	err = res.Body.Close()
	if err != nil {
		panic(err)
	}
}
