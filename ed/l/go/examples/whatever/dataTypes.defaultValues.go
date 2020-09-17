package main

import "fmt"

var (
	b bool
	s string
	i int
	f float32
	c complex64

	a   [2]int
	as  []int
	st  struct{ n int }
	st2 *struct{ n int }

	fn func()
	i2 interface{}
	m  map[int]int
	ch chan int
)

func main() {
	fmt.Printf("\n b: %+v", b) // b == nil # can not convert nil to ...
	fmt.Printf("\n s: %+v", s) // s == nil # can not convert nil to ...
	fmt.Printf("\n i: %+v", i) // i == nil # can not convert nil to ...
	fmt.Printf("\n f: %+v", f) // f == nil # can not convert nil to ...
	fmt.Printf("\n c: %+v", c) // c == nil # can not convert nil to ...
	fmt.Println()

	fmt.Printf("\n a: %+v", a)                              // a == nil # can not convert nil to []int
	fmt.Printf("\n as: %+v, is nil: %+v", as, as == nil)    // is nil
	fmt.Printf("\n st: %+v", st)                            // st == nil # can not convert nil to ...
	fmt.Printf("\n st2: %+v, is nil: %+v", st2, st2 == nil) // is nil
	fmt.Println()

	fmt.Printf("\n fn: %+v, is nil: %+v", fn, fn == nil)
	fmt.Printf("\n i2, is nil: %+v: %+v", i2, i2 == nil)
	fmt.Printf("\n m: %+v, is nil: %+v", m, m == nil)
	fmt.Printf("\n ch: %+v, is nil: %+v", ch, ch == nil)
	fmt.Println()
}

/*
 b: false
 s:
 i: 0
 f: 0
 c: (0+0i)

 a: [0 0]
 as: [], is nil: true
 st: {n:0}
 st2: <nil>, is nil: true

 fn: <nil>, is nil: true
 i2, is nil: <nil>: true
 m: map[], is nil: true
 ch: <nil>, is nil: true
*/
