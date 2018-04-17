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
	fmt.Printf("\n b: %+v", b)
	fmt.Printf("\n s: %+v", s)
	fmt.Printf("\n i: %+v", i)
	fmt.Printf("\n f: %+v", f)
	fmt.Printf("\n c: %+v", c)

	fmt.Printf("\n a: %+v", a)
	fmt.Printf("\n as: %+v, is nil: %+v", as, as == nil)
	fmt.Printf("\n st: %+v", st)
	fmt.Printf("\n st2: %+v", st2)

	fmt.Printf("\n fn: %+v", fn)
	fmt.Printf("\n i2: %+v", i2)
	fmt.Printf("\n m: %+v, is nil: %+v", m, m == nil)
	fmt.Printf("\n ch: %+v", ch)
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
 st2: <nil>

 fn: <nil>
 i2: <nil>
 m: map[], is nil: true
 ch: <nil>
*/
