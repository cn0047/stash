package main

import (
	"log"
)

const (
	name  = 1 << 0 // 1
	email = 1 << 1 // 2
	phone = 1 << 2 // 4
	ok    = name | email | phone
)

func main() {
	cur := 0

	cur |= name
	log.Printf("🔴 %#v \t %+v", cur, cur == ok)

	cur |= email
	log.Printf("🔴 %#v \t %+v", cur, cur == ok)

	cur |= phone
	log.Printf("🔴 %#v \t %+v", cur, cur == ok)
}
