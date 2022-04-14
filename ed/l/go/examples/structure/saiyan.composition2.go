package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func (p *Person) Introduce() {
	fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
	*Person
	Power int
}

func main() {
	goku := &Saiyan{
		Person: &Person{"Goku-In"},
		Power:  9001,
	}
	fmt.Println(goku.Name)
	fmt.Println(goku.Person.Name)
	goku.Introduce()
}
