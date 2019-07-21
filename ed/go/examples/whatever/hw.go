package main

import (
	"fmt"
)

func main() {
	name, ending := "World", "!"
	ending2 := `!!`
	fmt.Printf("Hello %s %s%s \n", name, ending, ending2)
}
