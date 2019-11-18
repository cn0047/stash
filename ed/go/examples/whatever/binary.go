/*

Number starts with 0       - base 8
Number never starts with 0 - base 10
Number starts with 0x      - base 16

&  - and
|  - or
^  - xor
&^ - bitclear (and not)
<< - left shift
>> - right shift

*/
package main

import (
	"log"
)

const (
	name  = 1 << 0 // 1 -> 0001
	email = 1 << 1 // 2 -> 0010
	phone = 1 << 2 // 4 -> 0100
	ok    = name | email | phone
)

func main() {
	one()
}

func one() {
	cur := 0

	cur |= name
	log.Printf("🔴 %#v \t %+v", cur, cur == ok)

	cur |= email
	log.Printf("🔴 %#v \t %+v", cur, cur == ok)

	cur |= phone
	log.Printf("🔴 %#v \t %+v", cur, cur == ok)
}
