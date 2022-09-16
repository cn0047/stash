package lib

import (
	"fmt"
	"io"
	"strconv"
)

str := strconv.FormatInt(int64, 10)                           // int64   -> str
str := strconv.FormatUint(uint64, 10)                         // uint64  -> str
str := strconv.FormatInt(int64(int32), 10)                    // int32   -> str
str := strconv.FormatInt(int64(intV), 10)                     // int     -> str
int, err := strconv.Atoi("-42")                               // str     -> int
int64, err := strconv.ParseInt(str, 10, 32); int32(int64)     // str     -> int32
int64, err := strconv.ParseInt(str, 10, 64);                  // str     -> int64
uint64, err := strconv.ParseUint(str, 10, 64)                 // str     -> uint64
s := fmt.Sprintf("%.0f", fl64); int32, err := strconv.Atoi(s) // float64 -> int32
↑; int64 = int64(int32)                                       // float64 -> int64
str := strconv.FormatBool(true)                               // bool    -> str

if !regexp.MustCompile(`^[\d]+$`).MatchString(ds) {}

// sortCharsInString represents string sort.
func sortCharsInString(str string) string {
	tmp := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		tmp = append(tmp, string(str[i]))
	}
	sort.Strings(tmp)
	return strings.Join(tmp, "")
}

func pow32(x, y int32) int32 {
	if y == 0 || x == 1 {
		return 1
	}
	if y == 1 {
		return x
	}

	var res = x
	for i := int32(1); i < y; i++ {
		res = res * x
	}

	return res
}

func pow64(x, y int64) int64 {
	if y == 0 || x == 1 {
		return 1
	}
	if y == 1 {
		return x
	}

	var res = x
	for i := int64(1); i < y; i++ {
		res = res * x
	}

	return res
}

// ReverseInt32 returns 321 for 123 etc.
func ReverseInt32(n int32) int32 {
	s := strconv.FormatInt(int64(n), 10)

	rs := ""
	for i := 0; i < len(s); i++ {
		rs = string(s[i]) + rs
	}

	r, _ := strconv.ParseInt(rs, 10, 32)

	return int32(r)
}

// SwapInIntSlice swaps i element with j element in array arr.
func SwapInIntSlice(arr []int, i int, j int) []int {
	arr[i], arr[j] = arr[j], arr[i]
	return arr
}

// @DEPRECATED.
func SwapInIntSlice(a []int, i int, j int) []int {
	v1 := a[i]
	v2 := a[j]

	a = append(a[:i], append([]int{v2}, a[i+1:]...)...)
	a = append(a[:j], append([]int{v1}, a[j+1:]...)...)

	return a
}

// SwapInStr swaps i char with j char in string s.
func SwapInStr(s string, i int, j int) string {
	c1 := s[i]
	c2 := s[j]

	s = s[:i] + string(c2) + s[i+1:]
	s = s[:j] + string(c1) + s[j+1:]

	return s
}

// LeftPadding adds char x to left part of string s up to final length n.
func LeftPadding(s string, x string, n int) string {
	if len(s) >= n {
		return s
	}

	m := n - len(s)
	res := ""
	for i := 0; i < m; i++ {
		res += x
	}

	res += s

	return res
}
