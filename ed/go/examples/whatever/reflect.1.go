package main

import (
	"fmt"
	"io"
	"reflect"
)

func main() {
	two()
}

func one() {
	var r io.Reader
	var w io.Writer
	w = r.(io.Writer)
	fmt.Printf("%#v", w)
}

func two() {
	a := 1
	fmt.Printf("\n%+v", reflect.TypeOf(a)) // int
	v := reflect.ValueOf(a)
	fmt.Printf("\n%+v", v) // 1
	fmt.Printf("\n%+v", v.Type()) // int
}
