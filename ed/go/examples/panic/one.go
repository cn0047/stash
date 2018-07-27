package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %+v\n", r)
		}
	}()

	Go()
}

func Go() {
	fmt.Println("Starting goroutine.")
	go func() {
		my()
	}()
}

func my() {
	panic(errors.New("some error"))
}
