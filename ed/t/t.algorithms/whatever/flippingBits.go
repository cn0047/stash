// see: https://www.hackerrank.com/challenges/flipping-bits/problem?h_l=interview&playlist_slugs%5B%5D%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D%5B%5D=miscellaneous
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(flippingBits(9) == int64(4294967286))
	fmt.Println(flippingBits(1) == int64(4294967294))
	fmt.Println(flippingBits(2147483647) == int64(2147483648))
	fmt.Println(flippingBits(4) == int64(4294967291))
	fmt.Println(flippingBits(123456) == int64(4294843839))
}

func flippingBits(n int64) int64 {
	b := strconv.FormatInt(n, 2)
	b = LeftPadding(b, "0", 32)

	res := ""
	for i := 0; i < len(b); i++ {
		c := "1"
		if b[i] == '1' {
			c = "0"
		}
		res += c
	}
	v, _ := strconv.ParseInt(res, 2, 64)

	return v
}

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
