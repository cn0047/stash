package main

import (
	"encoding/json"
	"fmt"
)

type Valuable interface {
	GetValue() string
}

func getValue(v Valuable) string {
	return v.GetValue()
}

type Foo struct {
	Foo string
}

func (f *Foo) GetValue() string {
	if f.Foo == "" {
		f.Foo = "Foo"
	}
	return f.Foo
}

func (f *Foo) GetFoo() string {
	return f.GetValue()
}

type Bar struct {
	Bar string
}

func (b *Bar) GetValue() string {
	if b.Bar == "" {
		b.Bar = "Bar"
	}
	return b.Bar
}

func (b *Bar) GetBar() string {
	return b.GetValue()
}

func main() {
	//f1()
	//f2()
	//f3()
	//f4()
	//f5()
	f6()
}

func f1() {
	f := Foo{}
	v := getValue(&f)
	fmt.Printf("[f1] %#v \n", v) // [f1] "Foo"
}

func f2() {
	f := Foo{}
	var v Valuable
	v = &f
	x, ok := v.(*Foo)
	fmt.Printf("[f2] v:%v | ok: %v x: %v \n", v.GetValue(), ok, x.GetValue()) // [f2] v:Foo | ok: true x: Foo
}

func f3() {
	f1 := Foo{Foo: "f3"}
	j, err := json.Marshal(f1)
	if err != nil {
	}

	var f2 Valuable
	err = json.Unmarshal(j, &f2)
	if err != nil {
	}

	v := getValue(f2) // panic: runtime error: invalid memory address or nil pointer dereference

	fmt.Printf("[f3] %#v \n", v)
}

func f4() {
	f1 := Foo{Foo: "f4"}
	j, err := json.Marshal(f1)
	if err != nil {
	}

	var f2 Foo
	err = json.Unmarshal(j, &f2)
	if err != nil {
	}

	fmt.Printf("[f4] %#v \n", getValue(&f2)) // [f4] "f4"
}

func f5() {
	f1 := Foo{Foo: "f5"}
	var if1 interface{}
	if1 = &f1

	j, err := json.Marshal(if1)
	if err != nil {
	}

	var f2 Foo
	err = json.Unmarshal(j, &f2)
	if err != nil {
	}

	fmt.Printf("[f5] %#v \n", getValue(&f2)) // [f5] "f5"
}

func f6() {
	f1 := Foo{Foo: "f6"}
	var v1 Valuable
	v1 = &f1

	j, err := json.Marshal(v1)
	if err != nil {
	}

	var f2 Foo
	err = json.Unmarshal(j, &f2)
	if err != nil {
	}

	fmt.Printf("[f6] %#v \n", getValue(&f2)) // [f6] "f6"
}
