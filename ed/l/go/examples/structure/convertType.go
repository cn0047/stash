package main

import (
	"fmt"
)

type Foo struct {
	ID   int
	Name string
}

type Bar struct {
	ID   int
	Name string
}

func main() {
	// two()
	one()
}

func two() {
	f := &Foo{ID: 1, Name: "foo"}
	b := Bar(*f)
	fmt.Printf("🔴 %#v", b) // main.Bar{ID:1, Name:"foo"}
}
