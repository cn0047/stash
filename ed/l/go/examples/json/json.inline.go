package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar Bar `json:"bar,inline"`
}

type Bar struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

func main() {
	v := Foo{Bar: Bar{Code: 2, Text: "bar"}}
	j, _ := json.Marshal(v)
	fmt.Printf("Result: %s \n", j)
}
