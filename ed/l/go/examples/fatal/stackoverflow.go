package main

func main() {
	var n int64 = 10 << 30
	// f1(n)
	f2(n)
}

func f2(n int64) {
}

func f1(n int64) {
	x := n + n
	f1(x) // runtime: goroutine stack exceeds 1000000000-byte limit
	// fatal error: stack overflow
}
