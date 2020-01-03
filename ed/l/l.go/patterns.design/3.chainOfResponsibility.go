package main

type Handler interface {
	Action()
}

type ConcreteHandlerA struct {
	Next Handler
}

func (c ConcreteHandlerA) Action() {
	println("A")
	if c.Next != nil {
		c.Next.Action()
	}
}

type ConcreteHandlerB struct {
	Next Handler
}

func (c ConcreteHandlerB) Action() {
	println("B")
}

func main() {
	a := ConcreteHandlerA{Next: ConcreteHandlerB{}}
	a.Action()
	// or
	ConcreteHandlerA{}.Action()
}
