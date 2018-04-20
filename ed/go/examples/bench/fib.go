package fib

var s = make([]int, 0)

func Fib(n int) int {
	if n < 2 {
		s = append(s, n) // for fun!)
		return n
	}

	return Fib(n-1) + Fib(n-2)
}
