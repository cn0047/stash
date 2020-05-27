package main

import (
	"fmt"
)

type G struct {
	x string
}

func (g G) f() string {
	return "it works!"
}

type P struct {
	k string
	v int
	G
}

func main() {
	//f1()
	f2()
}

func f2() {
	p := P{}
	p.k = "a"
	p.v = 2
	p.x = "b"
	fmt.Printf("[f2] %v | %v \n", p, p.f())   // [f2] {a 2 {b}} | it works!
	fmt.Printf("[f2] %v | %v \n", p, p.G.f()) // [f2] {a 2 {b}} | it works!
}

func f1() {
	p := P{k: "1", v: 2, G: G{x: "y"}}
	fmt.Printf("[f1] %v | %v \n", p, p.f()) // [f1] {1 2 {y}} | it works!
}
