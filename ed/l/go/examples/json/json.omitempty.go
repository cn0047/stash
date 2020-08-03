package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Msg string `json:"msg,omitempty"`
	Bar Bar    `json:"bar,omitempty"`
}

type Foo2 struct {
	Msg string `json:"msg,omitempty"`
	Bar *Bar   `json:"bar,omitempty"`
}

type Bar struct {
	Value string `json:"value,omitempty"`
}

func f1() {
	f := Foo{}
	j, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[f1] Result: %s \n", j) // [f1] Result: {"bar":{}}
}

func f2() {
	f := Foo2{}
	j, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[f2] Result: %s \n", j) // [f2] Result: {}
}

func f3() {
	f := Foo2{Bar: &Bar{}}
	j, err := json.Marshal(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[f3] Result: %s \n", j) //
}

func main() {
	//f1()
	//f2()
	f3()
}
