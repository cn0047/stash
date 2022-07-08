package main

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Reverse(s string) string {
	return ReverseV2(s)
}

func ReverseV1(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func ReverseV2(s string) string {
	if !utf8.ValidString(s) {
		fmt.Printf("input is not valid UTF-8")
		return ""
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func Test_Reverse(t *testing.T) {
	input := "abcd"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func Fuzz_Reverse(f *testing.F) {
	testCases := []string{"Hello, world", "!"}
	for _, tc := range testCases {
		f.Add(tc) // seeds
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}
