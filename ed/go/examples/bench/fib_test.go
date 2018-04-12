package fib

import "testing"

var fibTests = []struct {
	n        int // input
	expected int // expected result
}{
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}

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
