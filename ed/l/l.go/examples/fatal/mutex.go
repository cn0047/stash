package main

import (
	"fmt"
	"sync"
)

type PInt struct {
	val int
	sync.Mutex
}

func (v *PInt) Set(x int) {
	defer func() {
		r := recover()
		fmt.Printf("[RECOVERED] %#v \n", r)
	}()
	v.Lock()
	v.val = x
	v.Unlock()
	v.Unlock() // fatal error: sync: unlock of unlocked mutex
}

func main() {
	fmt.Println("Press enter.")
	fmt.Scanln()

	v := PInt{}
	v.Set(204)
	fmt.Printf("%#v \n", v)
}
