package main

import (
	"fmt"
)

type FooInterface interface {
	Bar()
}

type MyBoo struct {
}

func (m MyBoo) Bar() {
}

type MyFoo struct {
}

func (m MyFoo) Bar(str string) {
}

type MyBar struct {
}

func main() {
	f2()
}

func f2() {
	f := func(x interface{}) {
		y, ok := x.(interface{ Bar(str string) })
		fmt.Printf("[f2] ok: %v, res: %#v \n", ok, y)
	}
	f(MyBoo{})
	f(MyFoo{})
}

func f1() {
	var _ FooInterface = (*MyBoo)(nil)
	//var _ FooInterface = (*MyFoo)(nil) // cannot use (*MyFoo)(nil) (type *MyFoo) as type FooInterface
	//var _ FooInterface = (*MyBar)(nil) // cannot use (*MyBar)(nil) (type *MyBar) as type FooInterface
}
