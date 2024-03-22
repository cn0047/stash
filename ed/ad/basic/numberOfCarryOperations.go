// Something like manual math:
//
//   19
// +  1
//  ----
//   20
//
// carry = 1
//
package main

import (
	"fmt"
	"strconv"
)

func main() {
	println(numberOfCarryOperations(19, 1) == 1)
	println(numberOfCarryOperations(65, 55) == 2)
	println(numberOfCarryOperations(965, 55) == 3)
	println(numberOfCarryOperations(1000, 9) == 0)
	println(numberOfCarryOperations(1001, 9) == 1)
	println(numberOfCarryOperations(1, 999) == 3)
}

func numberOfCarryOperations(a int, b int) int {
	// return numberOfCarryOperationsV1(a, b)
	return numberOfCarryOperationsV2(a, b)
}

func numberOfCarryOperationsV2(a int, b int) int {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)

	// sa always contains longer string.
	if b > a {
		sa = sb
	}

	la := len(sa)
	lb := len(sb)

	res := 0
	carry := 0
	for i := 0; i < la; i++ {
		ca := string(sa[la-i-1])
		cb := "0"
		if lb-i-1 >= 0 {
			cb = string(sb[lb-i-1])
		}
		na, _ := strconv.Atoi(ca)
		nb, _ := strconv.Atoi(cb)

		if na + nb + carry >= 10 {
			res++
			carry = 1
		} else {
			carry = 0
		}
	}

	return res
}

// @deprecated
func numberOfCarryOperationsV1(a int, b int) int {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)
	la := len(sa)
	lb := len(sb)

	l := la
	if lb > l {
		l = lb
	}

	f := "%0" + strconv.Itoa(l) + "s"
	sa = fmt.Sprintf(f, sa)
	sb = fmt.Sprintf(f, sb)

	res := 0
	carry := 0
	for i := l - 1; i > 0; i-- {
		n1, _ := strconv.Atoi(string(sa[i]))
		n2, _ := strconv.Atoi(string(sb[i]))
		if n1+n2+carry >= 10 {
			res++
			carry = 1
		} else {
			carry = 0
		}
	}

	return res + carry
}
