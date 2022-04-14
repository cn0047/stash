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
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state += "iced;"
		iced <- cake
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

/*
{id:0 state:cooked;iced;}
{id:1 state:cooked;iced;}
{id:2 state:cooked;iced;}
{id:3 state:cooked;iced;}
{id:4 state:cooked;iced;}
{id:5 state:cooked;iced;}
{id:6 state:cooked;iced;}
{id:7 state:cooked;iced;}
{id:8 state:cooked;iced;}
{id:9 state:cooked;iced;}
*/
