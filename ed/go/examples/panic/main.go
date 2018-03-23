package main

import (
	"./lib"
	"fmt"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %+v\n", r)
		}
	}()

	lib.Go()
}
