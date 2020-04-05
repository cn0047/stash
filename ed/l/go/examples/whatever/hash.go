package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "my string"

	fmt.Printf("md5 = %x \n", md5.Sum([]byte(s)))

	sh1 := sha1.New()
	sh1.Write([]byte(s))
	fmt.Printf("sha1 = %x \n", sh1.Sum(nil))

	sh256 := sha256.New()
	sh256.Write([]byte(s))
	fmt.Printf("sha256 = %x \n", sh256.Sum(nil))
}
