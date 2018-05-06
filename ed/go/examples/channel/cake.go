package main

import (
	"fmt"
	"time"
)

type Cake struct {
	id    int
	state string
}

func baker(n int, cooked chan<- *Cake) {
	for i := 0; i < n; i++ {
		cake := new(Cake)
		cake.id = i
		cake.state = "cooked;"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state += "iced;"
		iced <- cake // icer never touches this cake again
	}
}

func done(iced <-chan *Cake) {
	for cake := range iced {
		fmt.Printf("%+v\n", *cake)
	}
}

func main() {
	cooked := make(chan *Cake)
	iced := make(chan *Cake)

	go baker(10, cooked)
	go icer(iced, cooked)
	go done(iced)

	time.Sleep(1 * time.Second)
}
