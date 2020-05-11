package main

import "fmt"

type Valuable interface {
	GetValue() string
}

func getValue(v Valuable) string {
	return v.GetValue()
}

type Foo struct {
	Foo string
}

func (f *Foo) GetFoo() string {
	return f.GetValue()
}

func (f *Foo) GetValue() string {
	return "Foo"
}

type Bar struct {
	Bar string
}

func (b *Bar) GetValue() string {
	return "Bar"
}

func (b *Bar) GetBar() string {
	return b.GetValue()
}

func main() {
	//f1()
	f2()
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
