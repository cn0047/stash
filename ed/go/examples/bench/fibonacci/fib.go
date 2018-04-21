package fibonacci

var pn int
var m = make(map[int]int)

func Fib(n int) int {
	if n < 2 {
		pn = n   // for fun!)
		m[n] = n // for fun!)
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
