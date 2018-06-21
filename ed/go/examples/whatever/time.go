package main

import (
	"fmt"
	"time"
)

func main() {
	f1()
}

func f1() {
	fmt.Printf("%+v ↔️ %+v \n", time.Now(), time.Now().Add(-24*time.Hour))
}
