package main

import (
	"fmt"
	"time"
)

func main() {
	//one()
	two()
}

func two() {
	c := make(chan struct{})
	go func() {
		<-c
		fmt.Println("ok")
	}()
	time.Sleep(1 * time.Second)
	close(c)
	fmt.Scanln()
}

func one() {
	c := make(chan struct{})
	go func() {
		<-c // hanging here
		fmt.Println("ok")
	}()
	fmt.Scanln()
}
