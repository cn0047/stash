package main

import (
	"fmt"
)

type Foo struct {
}

func (f Foo) F() {
	fmt.Println("foo")
}

type Bar struct {
}

func (b Bar) F() {
	fmt.Println("bar")
}

type X struct {
	Foo
	Bar
}

func main() {
	x := X{}
	x.F() // ambiguous selector x.F
}
