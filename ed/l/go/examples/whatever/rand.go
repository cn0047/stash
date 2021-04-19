package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// printRandomInts()
	printRandomStrings()
}

func randomStr(length int) string {
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}

func printRandomStrings() {
	for i := 1; i < 10; i++ {
		v := randomStr(16)
		fmt.Printf("random str val: %+v \n", v)
	}
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min) + min
}

func printRandomInts() {
	for i := 1; i < 10; i++ {
		v := randomInt(1, 2000)
		fmt.Printf("random int val: %+v \n", v)
	}
}
