package main

import (
	"fmt"
	"reflect"
)

type MyStr struct {
	Name string
}

func main() {
	m := MyStr{Name: "Foo"}
	e := reflect.ValueOf(&m).Elem()
	e.Field(0).SetString("Bar")

	fmt.Println(e.Type())
	fmt.Println(e.NumField())
	fmt.Println(m.Name)
}
