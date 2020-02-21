package main

import (
	"fmt"
)

type Foo struct {
	Bar
}

func (f Foo) F() {
	fmt.Println("foo")
}

type Bar struct {
}

func (b Bar) B() {
	fmt.Println("bar")
}

func (b Bar) F() {
	fmt.Println("foo from bar")
}

func main() {
	fmt.Println("1")
	f := Foo{}
	f.B()
	f.F()

	fmt.Println("2")
	b := Bar{}
	b.B()
	b.F()
}

/*
1
bar
foo
2
bar
foo from bar
*/
