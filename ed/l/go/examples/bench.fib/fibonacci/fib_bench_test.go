package fibonacci

import "testing"

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkAandB(b *testing.B) {
	b.Run("a", func(b *testing.B) {
		benchmarkFib(2, b)
	})

	b.Run("b", func(b *testing.B) {
		benchmarkFib(2, b)
	})
}

func BenchmarkC(b *testing.B) {
	benchmarkFib(1, b)
}

func BenchmarkD(b *testing.B) {
	benchmarkFib(1, b)
}

func BenchmarkFibComplete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(3)
	}
}
