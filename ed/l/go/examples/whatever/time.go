package main

import (
	"fmt"
	"time"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
	Usual      = "2006-01-02 15:04:05"
)

func main() {
	// f1()
	// f2()
	// f2b()
	// f3()
	// f4()
	// f5()
	// f6()
	// f7()
	AddDate()
}

func AddDate() {
	d, _ := time.Parse(Usual, "2020-11-27 10:58:59")
	years := 0
	months := 0
	days := 10
	fmt.Printf("[f8] \n\tberfor: %v \n\tafter:  %v \n", d, d.AddDate(years, months, -days))
}

func f7() {
	var t time.Time
	fmt.Printf("[f7|1] IsZero=%+v t=%v \n", t.IsZero(), t)
	fmt.Printf("[f7|2] now before t=%v \n", time.Now().Before(t))
}

func f6() {
	t1 := time.Unix(1581092731, 0).Unix()
	t2 := time.Now().Unix()
	d := t2 - t1
	fmt.Printf("diff in seconds between 2 unix timestamps = %#v \n", d)
}

func f5() {
	fmt.Printf("%s", time.Now().Format(time.Kitchen)) // 6:23PM
}

func f4() {
	t := time.Now()
	//stark-staging-events-1-2019-11-21-09-02-19-197cd971-83ef-4cee-af52-22f9c42ca30b
	// 2019/11/21/09/stark-staging-me-stream-1-2019-11-21-09-02-19-197cd971-83ef-4cee-af52-22f9c42ca30b

	s := fmt.Sprintf("%d-%02d-%02d_%d-%d-%d_%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	fmt.Printf("res: %d/%d/%d/%d/pref-%s-events-%s \n", t.Year(), t.Month(), t.Day(), t.Hour(), "dev", s)
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

	fmt.Printf("[f2|1] d1: %+v | d2: %+v \n", d1, d2) // 2018-08-01 00:00:00 +0000 UTC | d2: 2028-08-01 00:00:00 +0000 UTC

	fmt.Printf("[f2|2] %+v, %+v \n", d1.After(time.Now()), d2.After(time.Now()))   // false, true
	fmt.Printf("[f2|3] %+v, %+v \n", d1.Before(time.Now()), d2.Before(time.Now())) // true, false
	fmt.Printf("[f2|4] %+v, %+v \n", d1.Before(d2), d1.After(d2))                  // true, false
}

func f3() {
	dateFormat := "2006-01-02"
	d1, _ := time.Parse(dateFormat, "2018-08-31")

	duration := time.Since(d1)
	fmt.Printf("[f3] days: %+v \n", duration.Hours()/24) // [f3] days: 355.3300686416898
}
