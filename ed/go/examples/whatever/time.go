package main

import (
	"fmt"
	"time"
)

func main() {
	// f1()
	// f2()
	f2b()
	// f3()
}

func f1() {
	fmt.Printf("[f1] %+v ↔️ %+v \n", time.Now(), time.Now().Add(-24*time.Hour))
	// [f1] 2019-08-21 10:55:56.700883 +0300 EEST m=+0.000329841 ↔️ 2019-08-20 10:55:56.700883 +0300 EEST m=-86399.999670042
}

func f2b() {
	dateFormat := "2006-01-02 15:04:05"
	d, _ := time.Parse(dateFormat, "2019-08-21 10:53:25")
	fmt.Printf("%s", d)
}

func f2() {
	dateFormat := "2006-01-02"
	d1, _ := time.Parse(dateFormat, "2018-08-01")
	d2, _ := time.Parse(dateFormat, "2028-08-01")
	fmt.Printf(
		"[f2] 2️⃣ d1: %+v | d2: %+v | %+v, %+v \n",
		d1,
		d2,
		d1.After(time.Now()),
		d2.After(time.Now()),
	) // [f2] 2️⃣ d1: 2018-08-01 00:00:00 +0000 UTC | d2: 2028-08-01 00:00:00 +0000 UTC | false, true
}

func f3() {
	dateFormat := "2006-01-02"
	d1, _ := time.Parse(dateFormat, "2018-08-31")

	duration := time.Since(d1)
	fmt.Printf("[f3] days: %+v \n", duration.Hours()/24) // [f3] days: 355.3300686416898
}
