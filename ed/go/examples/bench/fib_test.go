package fib

import (
	"strconv"
	"testing"
)

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
	for i, tt := range fibTests {
		t.Run("Test case "+strconv.Itoa(i), func(t *testing.T) {
			actual := Fib(tt.n)
			if actual != tt.expected {
				t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
			}
		})
	}
}
