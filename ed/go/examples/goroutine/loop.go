package main

import "time"
import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 10; i++ {
		fmt.Printf("\n start: %d", i)
		go func(i int) {
			for j := 0; j < 10; j++ {
				fmt.Printf("\n \t %d| %d", i, j)
			}
		}(i)
	}

	time.Sleep(time.Second * 5)
}
