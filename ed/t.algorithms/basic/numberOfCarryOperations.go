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
	println(numberOfCarryOperations(65, 55) == 2)
	println(numberOfCarryOperations(965, 55) == 3)
	println(numberOfCarryOperations(1000, 9) == 0)
	println(numberOfCarryOperations(1001, 9) == 1)
}

func numberOfCarryOperations(a int, b int) int {
	sa := strconv.Itoa(a)
	sb := strconv.Itoa(b)
	na := len(sa)
	nb := len(sb)

	n := na
	if nb > n {
		n = nb
	}

	f := "%0" + strconv.Itoa(n) + "s"
	sa = fmt.Sprintf(f, sa)
	sb = fmt.Sprintf(f, sb)

	res := 0
	carry := 0
	for i := n - 1; i > 0; i-- {
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
