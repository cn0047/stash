package main

import (
	"fmt"
)

const (
	// ModeHorizontal prints matrix in natural horizontal mode (row after row).
	ModeHorizontal = 1
	// ModeVertical prints matrix in vertical mode (column after column).
	ModeVertical = 2
)

func main() {
	//printMatrix("123456789", 3, 3, ModeHorizontal)
	//printMatrix("12345678", 4, 2, ModeHorizontal)
	//printMatrix("12345678", 2, 4, ModeHorizontal)

	printMatrix("123456789", 3, 3, ModeVertical)
	//printMatrix("12345678", 4, 2, ModeVertical)
	//printMatrix("12345678", 2, 4, ModeVertical)
}

func printMatrix(matrix string, rows int, cols int, mode int) {
	if mode == ModeHorizontal {
		printInHorizontalMode(matrix, rows, cols)
		return
	}
	printInVerticalMode(matrix, rows, cols)
}

/*
	Input: 123456789, 3, 3
	Output:
	123
	456
	789

	Input: 12345678, 4, 2
	12
	34
	56
	78
*/
func printInHorizontalMode(matrix string, rows int, cols int) {
	for i := 0; i < len(matrix); i += cols {
		s := ""
		for j := 0; j < cols; j++ {
			s += string(matrix[i+j])
		}
		fmt.Printf("%s\n", s)
	}
}

/*
	Input: 123456789, 3, 3
	Output:
	147
	258
	369

	Input: 12345678, 4, 2
	15
	26
	37
	48
*/
func printInVerticalMode(matrix string, rows int, cols int) {
	for i := 0; i < rows; i++ {
		s := ""
		for j := i; j < len(matrix); j += rows {
			s += string(matrix[j])
		}
		fmt.Printf("%s\n", s)
	}
}
