package fibonacci

import (
	"strconv"
	"sync"
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
	{9, 34},
	{15, 610},
	{20, 6765},
	{27, 196418},
	{33, 3524578},
	{37, 24157817},
	{38, 39088169},
}

var numberForTests int

func init() {
	// InitNumberForTests can't be called here because here we don't have `t`.
	//InitNumberForTests()
}

func InitNumberForTests(t *testing.T) {
	// Mark function as a test helper function.
	t.Helper()

	numberForTests = 7
}

func TestFibSimple(t *testing.T) {
	InitNumberForTests(t)

	expected := 13
	actual := Fib(numberForTests)

	if actual != expected {
		t.Errorf("Got: %v, want: %v", actual, expected)
	}
}

func TestFib(t *testing.T) {
	for i, tt := range fibTests {
		t.Run("Test case "+strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			actual := Fib(tt.n)
			if actual != tt.expected {
				t.Errorf("Fib(%d): expected %d, actual %d", tt.n, tt.expected, actual)
			}
		})
	}
}

func TestFibInGoroutine(t *testing.T) {
	m := make(map[string]int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		m["case-1"] = Fib(5)
	}()
	go func() {
		defer wg.Done()
		m["case-2"] = Fib(9)
	}()

	wg.Wait()

	if m["case-1"] != 5 {
		t.Errorf("Got: %v, want: %v", m["case-1"], 5)
	}
	if m["case-2"] != 34 {
		t.Errorf("Got: %v, want: %v", m["case-2"], 34)
	}
}
