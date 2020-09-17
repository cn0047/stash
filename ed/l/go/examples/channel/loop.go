package main

import (
	"fmt"
	"time"
)

func main() {
	breakNotStopSelect()
}

func breakNotStopSelect() {
	c := make(chan int, 1)
	c <- 1

	go func() {
		for {
			select {
			case c <- 2:
				print("w")
				break // won't work
			default:
				time.Sleep(100 * time.Millisecond)
				print(".")
			}
		}
		print("|")
	}()

	go func() {
		time.Sleep(1 * time.Second)
		<-c
		print("r")
	}()

	fmt.Scanln()
}
