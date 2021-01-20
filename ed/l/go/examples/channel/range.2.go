package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		for m := range ch {
			println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2"
	ch <- "cmd.3"

	fmt.Scanln()
}

/*
processed: cmd.1
processed: cmd.2
processed: cmd.3
*/
