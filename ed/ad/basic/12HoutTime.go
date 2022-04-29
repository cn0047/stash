package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(convertTo24HoutTime("12:01:00AM") == "00:01:00")
	fmt.Println(convertTo24HoutTime("01:01:00AM") == "01:01:00")
	fmt.Println(convertTo24HoutTime("12:01:00PM") == "12:01:00")
	fmt.Println(convertTo24HoutTime("01:01:00PM") == "13:01:00")
	fmt.Println(convertTo24HoutTime("11:01:00PM") == "23:01:00")
	fmt.Println(convertTo24HoutTime("07:05:45PM") == "19:05:45")
}

// convertTo24HoutTime converts input 12 hour time string to 24 hour time.
// IMPORTANT: input string must follows format: hh:mm:ss:(AM|PM)
func convertTo24HoutTime(s string) string {
	hh, _ := strconv.Atoi(s[:2])
	mm := s[3:5]
	ss := s[6:8]
	t := s[8:]

	if t == "AM" && hh == 12 {
		hh = 0
	}
	if t == "PM" && hh != 12 {
		hh += 12
	}

	r := fmt.Sprintf("%02d:%s:%s", hh, mm, ss)

	return r
}
