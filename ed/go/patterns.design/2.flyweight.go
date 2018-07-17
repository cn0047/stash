package main

import "fmt"

type Flyweight struct {
	Name string
}

type FlyweightFactory struct {
	pool map[string]Flyweight
}

func NewFlyweightFactory() FlyweightFactory {
	return FlyweightFactory{pool: make(map[string]Flyweight)}
}

func (f FlyweightFactory) GetFlyweight(str string) Flyweight {
	flyweight, exists := f.pool[str]
	if exists == true {
		fmt.Println("use existing: " + str)
		return flyweight
	}

	fmt.Println("new: " + str)
	fw := Flyweight{str}
	f.pool[str] = fw

	return fw
}

func main() {
	f := NewFlyweightFactory()
	f.GetFlyweight("foo")
	f.GetFlyweight("foo")
	f.GetFlyweight("foo")
	f.GetFlyweight("bar")
}
