package main

import (
	"fmt"
	"testing"
)

func x(t *testing.T, n int) {
	if n == -2 {
		t.Fail()
	}
	if n == -1 {
		t.Fatal("[x] Fatal.")
	}
	if n == 0 {
		t.Skip("[x] Skip.")
	}
}

func TestOne(t *testing.T) {
	t.Run("-2", func(st *testing.T) {
		x(st, -2)
		fmt.Printf("[TestOne] -2")
	})

	t.Run("-1", func(st *testing.T) {
		return
		x(st, -1)
		fmt.Printf("[TestOne] -1")
	})

	t.Run("0", func(st *testing.T) {
		x(st, 0)
		fmt.Printf("[TestOne] 0")
	})

	t.Run("1", func(st *testing.T) {
		x(st, 1)
		fmt.Printf("[TestOne] 1")
	})
}

// go test ed/go/examples/whatever/test_test.go
