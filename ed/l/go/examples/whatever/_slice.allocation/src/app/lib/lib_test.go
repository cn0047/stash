package lib

import "testing"

func BenchmarkF1(b *testing.B) {
	F1(100)
}

func BenchmarkF2(b *testing.B) {
	F2(100)
}
