package fib

import "testing"

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib1(b *testing.B) {
	benchmarkFib(1, b)
}

func BenchmarkFib5(b *testing.B) {
	benchmarkFib(5, b)
}

func BenchmarkFib10and20(b *testing.B) {
	b.Run("10", func(b *testing.B) {
		benchmarkFib(10, b)
	})

	b.Run("20", func(b *testing.B) {
		benchmarkFib(20, b)
	})
}

func BenchmarkFibComplete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(10)
	}
}
