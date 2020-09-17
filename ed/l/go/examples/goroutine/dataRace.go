package main

var (
	n = 0
)

func main() {
	for i := 0; i < 5; i++ {
		go func(i int) {
			println(i)
			n++ // here
		}(i)
	}
	for {
		if n == 5 {
			break
		}
	}
}

/*
1
3
4
0
2
*/
