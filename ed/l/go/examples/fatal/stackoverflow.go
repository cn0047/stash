package main

func main() {
	var n int64 = 10 << 30
	f(n)
}

func f(n int64) {
	x := n + n
	f(x) // runtime: goroutine stack exceeds 1000000000-byte limit
	// fatal error: stack overflow
}
