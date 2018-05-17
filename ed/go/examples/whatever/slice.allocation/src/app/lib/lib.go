package lib

func F1(n int) {
	s := make([]int, 0)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
}

func F2(n int) {
	s := make([]int, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
}
