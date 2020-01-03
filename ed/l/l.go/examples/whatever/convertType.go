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
	two()
}

func two() {
	f := &Foo{ID: 1, Name: "foo"}
	b := Bar(*f)
	fmt.Printf("🔴 %#v", b)
}

func one() {
	a := 5.0
	b := int(a)
	c := uint(b)
	fmt.Printf("\n a = %T, b = %T, c = %T", a, b, c)
}
