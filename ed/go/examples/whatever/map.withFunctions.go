package main

import (
	"fmt"
)

var (
	Funcs = map[string]MyFunc{
		"foo": Foo,
		"bar": Bar,
	}
)

type MyFunc func(s string)

func main() {
	f := "bar"
	Funcs[f]("test")
}

func Foo(s string) {
	fmt.Println("Foo:", s)
}

func Bar(s string) {
	fmt.Println("Bar:", s)
}
