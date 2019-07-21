package main

import (
	"fmt"
)

type MyErr struct {
}

func (m MyErr) Error() string {
	return "MyError"
}

func main() {
	err := f()
	fmt.Printf("🔴 %+v \n", err == nil) // err is NOT nil
}

func f() error {
	var e *MyErr
	return e
}
