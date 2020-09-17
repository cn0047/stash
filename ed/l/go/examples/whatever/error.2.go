package main

import (
	"errors"
	"fmt"
)

var (
	ERR_2 = fmt.Errorf("f2 error.")
)

func main() {
	err := f1()
	fmt.Printf("1) %v \n", err)                                   // f1 error to call next func: f2 error.
	fmt.Printf("2) error is ERR_2: %v \n", errors.Is(err, ERR_2)) // true
	fmt.Printf("3) %v \n", errors.Unwrap(err))                    // f2 error.
}

func f1() error {
	err := f2()
	if err != nil {
		return fmt.Errorf("f1 error to call next func: %w", err)
	}
	return nil
}

func f2() error {
	return ERR_2
}
