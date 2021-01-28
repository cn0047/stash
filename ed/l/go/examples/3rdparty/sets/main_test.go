package main

import (
	"fmt"
	"testing"
)

func assert(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Got: %#v, want: %#v", actual, expected)
	}
}

func Test_findSubSetCommand(ts *testing.T) { // test suite
	ts.Run("invalid", func(t *testing.T) {
		s := "[ SUM a b "
		r := findSubSetCommand(s)

		assert(t, 0, len(r))
	})

	ts.Run("simple case", func(t *testing.T) {
		s := "[ SUM a.txt b.txt ]"
		r := findSubSetCommand(s)

		assert(t, s, r[0][0])
	})

	ts.Run("with subsets", func(t *testing.T) {
		s := "[ SUM [ DIF a.txt b.txt ] [ INT b.txt c.txt ] ]"
		r := findSubSetCommand(s)

		assert(t, "[ DIF a.txt b.txt ]", r[0][0])
		assert(t, "[ INT b.txt c.txt ]", r[1][0])
	})
}

func Test_loadIntoSlice(ts *testing.T) {
	ts.Run("simple case", func(t *testing.T) {
		s, err := loadIntoSlice("a.txt")

		assert(t, "[1 2 3]", fmt.Sprintf("%v", s))
		assert(t, nil, err)
	})
}

func Test_loadIntoMap(ts *testing.T) {
	ts.Run("simple case", func(t *testing.T) {
		m, err := loadIntoMap("a.txt")

		assert(t, "map[1:{} 2:{} 3:{}]", fmt.Sprintf("%v", m))
		assert(t, nil, err)
	})
}

func Test_intsc(ts *testing.T) {
	ts.Run("simple case", func(t *testing.T) {
		r, err := intsc("a.txt", "b.txt")

		assert(t, "[2 3]", fmt.Sprintf("%v", r))
		assert(t, nil, err)
	})
}

func Test_sdiff(ts *testing.T) {
	ts.Run("simple case", func(t *testing.T) {
		r, err := sdiff("a.txt", "b.txt")

		assert(t, "[1 4]", fmt.Sprintf("%v", r))
		assert(t, nil, err)
	})
}

func Test_sum(ts *testing.T) {
	ts.Run("simple case", func(t *testing.T) {
		r, err := sum("a.txt", "b.txt")

		assert(t, "[1 2 3 4]", fmt.Sprintf("%v", r))
		assert(t, nil, err)
	})
}
