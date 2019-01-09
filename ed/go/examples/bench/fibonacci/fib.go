package fibonacci

var s = make([]int, 0)

func Fib(n int) int {
	s = append(s, n) // for fun!)
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
