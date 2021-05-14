package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//f1()
	//f2()
	//f3()
	// sliceOfStructs()
	sliceOfStructsPointers()
}

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

type Bars struct {
	Data []Bar `json:"data,omitempty"`
}

type BarsPointers struct {
	Data []*Bar `json:"data,omitempty"`
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
	fmt.Printf("[f3] Result: %s \n", j) // [f3] Result: {"bar":{}}
}

func sliceOfStructs() {
	b := Bars{}
	b.Data = []Bar{
		{Value: "1"},
		{Value: "2"},
		{Value: ""},
	}
	j, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[sliceOfStructs] Result: %s \n", j) // Result: {"data":[{"value":"1"},{"value":"2"},{}]}
}

func sliceOfStructsPointers() {
	b := BarsPointers{}
	b.Data = []*Bar{
		{Value: "1"},
		{Value: "2"},
		{Value: ""},
	}
	j, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[sliceOfStructsPointers] Result: %s \n", j) // [sliceOfStructsPointers] Result: {"data":[{"value":"1"},{"value":"2"},{}]}
}
