package main

import (
	"fmt"
)

type MyNewType struct {
	N int
}

func main() {
	mnt := new(MyNewType)
	fmt.Printf("%+v", mnt)
}
