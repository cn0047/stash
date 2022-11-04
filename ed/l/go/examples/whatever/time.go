package main

import (
	"fmt"
	"math"
	"strconv"
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
	// timeMihusSomeHours()
	// datesComparison()
	// parseDate()
	// calsSinceDuration()
	// printTimeParts()
	// formatToKitchen()
	// unixTimestampsDiff()
	// IsZero()
	// AddDate()
	// AfterFunc()
	// convertLocation()
	// printTimezoneInfo()
	generateHoursRange()
}

func AfterFunc() {
	time.AfterFunc(1*time.Second, func() {
		fmt.Printf("[AfterFunc] 2\n")
	})
	fmt.Printf("[AfterFunc] 1\n")
	_, _ = fmt.Scanln()
}

func AddDate() {
	d, _ := time.Parse(Usual, "2020-11-27 10:58:59")
	years := 0
	months := 0
	days := 10
	fmt.Printf("[f8] \n\tberfor: %v \n\tafter:  %v \n", d, d.AddDate(years, months, -days))
}

func IsZero() {
	var t time.Time
	fmt.Printf("IsZero=%+v; t=%v; \n", t.IsZero(), t) // IsZero=true; t=0001-01-01 00:00:00 +0000 UTC;
}

func unixTimestampsDiff() {
	t1 := time.Unix(1581092731, 0).Unix()
	t2 := time.Now().Unix()
	d := t2 - t1
	fmt.Printf("unixTimestampsDiff = %#v seconds \n", d) // unixTimestampsDiff = 86297586 seconds
}

func formatToKitchen() {
	fmt.Printf("%s", time.Now().Format(time.Kitchen)) // 6:23PM
}

func printTimeParts() {
	t := time.Now()
	//stark-staging-events-1-2019-11-21-09-02-19-197cd971-83ef-4cee-af52-22f9c42ca30b
	// 2019/11/21/09/stark-staging-me-stream-1-2019-11-21-09-02-19-197cd971-83ef-4cee-af52-22f9c42ca30b

	s := fmt.Sprintf("%d-%02d-%02d_%d-%d-%d_%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
	fmt.Printf("res: %d/%d/%d/%d/pref-%s-events-%s \n", t.Year(), t.Month(), t.Day(), t.Hour(), "dev", s)
	// res: 2022/11/2/13/pref-dev-events-2022-11-02_13-2-3_798528000
}

func timeMihusSomeHours() {
	fmt.Printf("now: %+v \nres: %+v \n", time.Now(), time.Now().Add(-24*time.Hour))
	// now: 2022-11-02 15:40:38.047317 +0100 CET m=+0.000145918
	// res: 2022-11-01 15:40:38.047317 +0100 CET m=-86399.999853898
}

func parseDate() {
	dateFormat := "2006-01-02 15:04:05"
	d, _ := time.Parse(dateFormat, "2019-08-21 10:53:25")
	fmt.Printf("res: %s", d) // res: 2019-08-21 10:53:25 +0000 UTC
}

func datesComparison() {
	dateFormat := "2006-01-02"
	d1, _ := time.Parse(dateFormat, "2018-08-01")
	d2, _ := time.Parse(dateFormat, "2028-08-01")

	fmt.Printf("[1] d1: %+v | d2: %+v \n", d1, d2) // 2018-08-01 00:00:00 +0000 UTC | d2: 2028-08-01 00:00:00 +0000 UTC

	fmt.Printf("[2] %+v, %+v \n", d1.After(time.Now()), d2.After(time.Now()))   // false, true
	fmt.Printf("[3] %+v, %+v \n", d1.Before(time.Now()), d2.Before(time.Now())) // true, false
	fmt.Printf("[4] %+v, %+v \n", d1.Before(d2), d1.After(d2))                  // true, false
	fmt.Printf("[5] %+v, %+v \n", time.Now().Before(d1), time.Now().Before(d2)) // false, true
}

func calsSinceDuration() {
	dateFormat := "2006-01-02"
	d1, _ := time.Parse(dateFormat, "2018-08-31")

	duration := time.Since(d1)
	fmt.Printf("days: %+v \n", duration.Hours()/24) // [f3] days: 355.3300686416898
}

func printTimezoneInfo() {
	t := time.Now()
	z, o := t.Zone()
	l := t.Location()
	fmt.Printf("zone: %+v; offset in seconds: %+v; location: %+v \n", z, o, l) // zone: CET; offset in seconds: 3600; location: Local
}

func convertLocation() {
	l, _ := time.LoadLocation("Europe/Kyiv")
	t := time.Now()

	fmt.Printf(" now: %+v; \n UTC: %+v; \nKyiv: %+v \n\n", t, t.In(time.UTC), t.In(l))
	//  now: 2022-11-02 18:00:38.425154 +0100 CET m=+0.001216652; in Warsaw
	//  UTC: 2022-11-02 17:00:38.425154 +0000 UTC;
	// Kyiv: 2022-11-02 19:00:38.425154 +0200 EET

	t1 := time.Date(2020, 1, 19, 9, 0, 0, 0, time.UTC)
	t2 := time.Date(2020, 1, 19, 11, 0, 0, 0, l)

	fmt.Printf("t1: %+v; in UTC: %+v; in Kyiv: %+v \n", t1, t1.In(time.UTC), t1.In(l))
	fmt.Printf("t2: %+v; in UTC: %+v; in Kyiv: %+v \n", t2, t2.In(time.UTC), t2.In(l))
	fmt.Printf("equal: %+v \n", t1.Equal(t2))
	// t1: 2020-01-19 09:00:00 +0000 UTC; in UTC: 2020-01-19 09:00:00 +0000 UTC; in Kyiv: 2020-01-19 11:00:00 +0200 EET
	// t2: 2020-01-19 11:00:00 +0200 EET; in UTC: 2020-01-19 09:00:00 +0000 UTC; in Kyiv: 2020-01-19 11:00:00 +0200 EET
	// equal: true
}

func generateHoursRange() {
	generateRangeForHours := 7
	t1 := time.Now()
	t2 := time.Now().Add(time.Duration(generateRangeForHours) * time.Hour)
	diffInHoursFloat := math.Round(t2.Sub(t1).Hours())
	diffInHours, _ := strconv.Atoi(fmt.Sprintf("%.0f", diffInHoursFloat))

	for i := 0; i < diffInHours; i++ {
		t1 = t1.Add(1 * time.Hour)
		fmt.Printf("%v) %+v:00 \n", i, t1.Hour())
	}
}
