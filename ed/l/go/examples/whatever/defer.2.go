package main

import (
	"fmt"
)

type Struct struct {
	id string
}

func (s Struct) print1() {
	fmt.Println(s.id)
}
func (s *Struct) print2() {
	fmt.Println(s.id)
}

func main() {
	s1 := Struct{id: "s1:foo"}
	defer s1.print1()
	s1.id = "s1:bar"

	s2 := &Struct{id: "s2:foo"}
	defer s2.print2()
	s2.id = "s2:bar"
	// Result:
	// s2:bar
	// s1:foo
}
