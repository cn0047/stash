package main

import (
	"fmt"
	"time"
)

func main() {
	go f(nil)
	time.Sleep(time.Second * 5)
	fmt.Println("exit")
}

func f(ch chan int) {
	fmt.Println("block")
	fmt.Printf("%+v", <-ch)
	fmt.Println("unblock")
}

/*
block
exit
*/
