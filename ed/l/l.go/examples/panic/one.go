package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %+v\n", r)
		}
	}()

	Go()
	time.Sleep(time.Second * 2)
}

func Go() {
	fmt.Println("Starting goroutine.")
	go func() {
		my()
	}()
}

func my() {
	fmt.Println("Panicking...")
	panic(errors.New("some error"))
}
