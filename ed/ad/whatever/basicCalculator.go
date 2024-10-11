package main

import (
	"fmt"
)

func main() {
	r := 0
	r = basicCalculator1("3+2*2")     // 7
	r = basicCalculator1(" 3+5 / 2 ") // 5

	// r = basicCalculator2(" 2-1 + 2 ")           // 3
	// r = basicCalculator2("(1+(4+5+2)-3)+(6+8)") // 23
	// r = basicCalculator2("-(2 + 3)")            // -5
	// r = basicCalculator2("(130-10) - (5+15)")   // 100

	fmt.Printf("\n===\n%v", r)
}

// basicCalculator1 holds implementation for basic calculator.
// https://leetcode.com/problems/basic-calculator-ii/description
func basicCalculator1(s string) (res int) {
	s = s + "=" // to have one loop, and to perform action after last digit read for last number in string.
	n := len(s)
	stack := []int{}
	action := byte('+')
	val := 0

	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			continue
		} else if '0' <= s[i] && s[i] <= '9' {
			val = val*10 + int(s[i]-'0') // multiply to 10 to grow number by 10, 100, 1000, etc.
		} else {
			// Current s[i] means that finished read for number and it's time to fulfil previous action.
			switch action {
			case '+':
				stack = append(stack, val)
			case '-':
				stack = append(stack, -val)
			case '*':
				stack[len(stack)-1] *= val
			case '/':
				stack[len(stack)-1] /= val
			case '=':
				// Last character in input string, it's like termination.
			}
			action = s[i] // assign next action, which is defined by current s[i].
			val = 0 // assign default value, to start read new number.
		}
	}

	for _, val := range stack {
		res += val
	}

	return res
}

// basicCalculator2 holds implementation for basic calculator with parentheses.
// @see *: https://leetcode.com/problems/basic-calculator
func basicCalculator2(s string) (res int) {
	n := len(s)
	stack := []int{}
	sign := 1

	for i := 0; i < n; i++ {
		switch s[i] {
		case ' ':
		case '+':
			sign = 1
		case '-':
			sign = -1
		case '(':
			stack = append(stack, res)  // add current value into stack.
			stack = append(stack, sign) // add current sign into stack.
			res, sign = 0, 1            // assign default values.
		case ')':
			res *= stack[len(stack)-1]   // multiply to last sign.
			stack = stack[:len(stack)-1] // remove first element from stack (shift).
			res += stack[len(stack)-1]   // add last value to accumulative res.
			stack = stack[:len(stack)-1] // remove first element from stack (shift).
		default:
			val := 0
			j := i
			for ; j < n && '0' <= s[j] && s[j] <= '9'; j++ { // iteratively read the number.
				val = val*10 + int(s[j]-'0') // multiply to 10 to grow number by 10, 100, 1000, etc.
			}
			i = j - 1
			res += sign * val
		}
	}

	return res
}
