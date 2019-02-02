package main

import "fmt"

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
  f := Foo{}
  f.F()
  f.B()
}
