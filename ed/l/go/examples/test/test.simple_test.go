package main

import (
	"testing"
)

func x(s string) string {
	return s
}

func Test_X(ts *testing.T) {
	ts.Run("Simple case 1", func(t *testing.T) {
		actual := x("foo")
		expected := "foo"

		if actual != expected {
			t.Errorf("Got: %#v, want: %#v", actual, expected)
		}
	})
}
