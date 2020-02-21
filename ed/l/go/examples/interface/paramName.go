package main

type A struct{}

func (this A) Do(argument string) {}

type B struct{}

func (this B) Do(arg string) {}

type I interface {
	Do(a string)
}

func main() {
	b := B{}
	f(b)
}

func f(element I) {
	element.Do("test") // ok
}
