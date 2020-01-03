package main

func main() {
	var f func(a [1000]int64)
	f = func(a [1000]int64) {
		f(a) // fatal error: stack overflow
	}
	f([1000]int64{})
}
