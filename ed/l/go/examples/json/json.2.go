package main

import (
	"encoding/json"
	"fmt"
)

type Req1 struct {
	Val string `json:"val"`
	Hidden string `json:"-"`
}

type Req2 struct {
	Val string
	Hidden string
}

func f1() {
	p := []byte(`{"val": "yes", "hidden": "no"}`)
	r := Req1{}
	err := json.Unmarshal(p, &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", r)
}

func f2() {
	p := []byte(`{"val": "yes", "hidden": "no"}`)
	r := Req2{}
	err := json.Unmarshal(p, &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v \n", r)
}

func main() {
	f2()
}
