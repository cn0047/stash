package main

import (
	"fmt"
)

func main() {
	var x Animal = Dog{"Rosie"}

	if x, ok := x.(Human); ok { // in case of fail x contains ZERO VALUE of type (Human)
		fmt.Println(x.name, "doesn't want to be treated like dog.")
	} else {
		fmt.Println(x.Say())
	}
}

type Animal interface {
	Say() string
}

type Dog struct {
	name string
}

func (d Dog) Say() string {
	return fmt.Sprintf("%v barks", d.name)
}

type Human struct {
	name string
}

func (h Human) Say() string {
	return fmt.Sprintf("%v speaks", h.name)
}
