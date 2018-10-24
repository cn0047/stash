package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestData1 struct {
	Code   string `json:"code"`
	Status string `json:"status"`
}

type RequestData2 struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/get", get1)
	http.HandleFunc("/post", post2)
	http.ListenAndServe(":8080", nil)
}

// @see curl 'http://localhost:8080/get?f=foo&b=bar'
func get1(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	fmt.Printf("Query: %+v", query)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Look into console.`))
}

func post1(w http.ResponseWriter, r *http.Request) {
	body1, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	rd1 := RequestData1{}
	err = json.Unmarshal(body1, &rd1)
	if err != nil {
		panic(err)
	}

	body2, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	rd2 := RequestData2{}
	err = json.Unmarshal(body2, &rd2)
	if err != nil {
		panic(err) // 2018/07/19 17:56:01 http: panic serving [::1]:54581: unexpected end of JSON input
	}

	fmt.Printf("rd1: %+v \nrd2: %+v", rd1, rd2)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Look into console.`))
}

// @see curl -X POST 'http://localhost:8080/post' -H 'Content-Type: application/json' -d '{"code":"200", "status": "OK", "message": "200 OK"}'
func post2(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	rd1 := RequestData1{}
	err = json.Unmarshal(body, &rd1)
	if err != nil {
		panic(err)
	}

	rd2 := RequestData2{}
	err = json.Unmarshal(body, &rd2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("rd1: %+v \nrd2: %+v", rd1, rd2)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`Look into console.`))
}
