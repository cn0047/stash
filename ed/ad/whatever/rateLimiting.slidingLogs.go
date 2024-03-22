// Sliding logs holds requests with time in queue (memory eager).
// Worst case for request = O(n).
package main

import (
	"fmt"
	"time"
)

var (
	log = map[string][]time.Time{}
)

func main() {
	fmt.Println(request("user-1"))
	time.Sleep(20 * time.Second)
	fmt.Println(request("user-1"))
	time.Sleep(20 * time.Second)
	fmt.Println(request("user-1"))
	time.Sleep(20 * time.Second)
	fmt.Println(request("user-1"))
	time.Sleep(5 * time.Second)
	fmt.Println(request("user-1"))
}

// request allows 3 requests per 1 minute.
func request(user string) string {
	_, ok := log[user]
	if !ok {
		// Init logs for user.
		t := time.Now()
		log[user] = []time.Time{t} // init.
		return fmt.Sprintf("ok - 1 %s", t)
	}

	// Add new log.
	nt := time.Now()
	log[user] = append(log[user], nt) // add time.

	// Delete old logs.
	for len(log[user]) > 0 {
		t := log[user][0]
		diff := time.Since(t)
		if diff.Minutes() >= 1 { // IMPORTANT: 1 minute is threshold.
			// Delete old log.
			log[user] = log[user][1:]
			fmt.Println("del:", t)
		} else {
			// No use to continue loop, next logs even more fresh.
			break
		}
	}

	// Throttle.
	if len(log[user]) > 3 {
		return fmt.Sprintf("sorry %s", nt)
	}

	return fmt.Sprintf("ok - 2 %s", nt)
}
