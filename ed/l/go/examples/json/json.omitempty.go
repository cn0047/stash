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
	// sliceOfStructsPointers()
	// boolValue()
	boolSlice()
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

type BoolStruct struct {
	Value bool `json:"value,omitempty`
}

type BoolBag struct {
	Data []BoolStruct `json:"data,omitempty`
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
	fmt.Printf("[sliceOfStructs] Result: %s \n", j) // {"data":[{"value":"1"},{"value":"2"},{}]}
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
	fmt.Printf("[sliceOfStructsPointers] Result: %s \n", j) // {"data":[{"value":"1"},{"value":"2"},{}]}
}

func boolValue() {
	s1 := BoolStruct{Value: true}
	j1, err := json.Marshal(s1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[boolValue] Result 1: %s \n", j1) // {"Value":true}

	s2 := BoolStruct{Value: false}
	j2, err := json.Marshal(s2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[boolValue] Result 2: %s \n", j2) // {"Value":false}
}

func boolSlice() {
	d := BoolBag{Data: []BoolStruct{{Value: true}, {Value: false}}}
	j, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[boolSlice] Result: %s \n", j) // {"Data":[{"Value":true},{"Value":false}]}
}
