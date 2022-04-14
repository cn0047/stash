package main

import (
	"fmt"
)

func main() {
	var v AliveBeing = Animal{name: "Rex"} // Human - NOT ok, <> speaks
	//var v AliveBeing = Human{name: "Bob"} // Human - ok, <Bob> speaks

	x, ok := v.(Human) // in case of fail x contains ZERO VALUE of type (Human)
	if ok {
		fmt.Printf("Human - ok, %s \n", x.Say())
	} else {
		fmt.Printf("Human - NOT ok, %s \n", x.Say())
	}
}

type AliveBeing interface {
	Say() string
}

type Animal struct {
	name string
}

func (a Animal) Say() string {
	return fmt.Sprintf("<%v> barks", a.name)
}

type Human struct {
	name string
}

func (h Human) Say() string {
	return fmt.Sprintf("<%v> speaks", h.name)
}
