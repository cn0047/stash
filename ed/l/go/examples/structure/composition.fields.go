package main

import (
	"fmt"
)

type G struct {
	x string
}

func (g G) f() {
	fmt.Printf("it works \n")
}

type P struct {
	k string
	v int
	G
}

func main() {
	p := P{k: "1", v: 2, G: G{x: "y"}}
	fmt.Printf("%v \n", p)
	p.f() // it works
}
