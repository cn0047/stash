package main

var (
	n = 0
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			println(i)
			n++
		}(i)
	}
	for {
		if n == 5 {
			break
		}
	}
}
