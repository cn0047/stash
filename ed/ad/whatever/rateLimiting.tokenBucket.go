// Token bucket holds for requests only tokens and last access time.
// Worst case for request = O(1).

package main

import (
	"fmt"
	"time"
)

type AccessRate struct {
	Time   time.Time
	Tokens int
}

var (
	store = map[string]AccessRate{}
)

func main() {
	timeToSleep := 100 * time.Millisecond

	fmt.Println(request("user-1"))
	time.Sleep(timeToSleep)
	fmt.Println(request("user-1"))
	time.Sleep(timeToSleep)
	fmt.Println(request("user-1"))
	time.Sleep(timeToSleep)
	fmt.Println(request("user-1"))
}

// request allows 3 requests per 1 second.
func request(user string) string {
	accessRate, ok := store[user]
	if !ok {
		// Init user's tokens.
		store[user] = AccessRate{Time: time.Now(), Tokens: 3 - 1}
		return "ok - 1"
	}

	diff := time.Since(accessRate.Time)
	if diff.Seconds() > 1 { // IMPORTANT: 1 second is threshold.
		// Refresh tokens.
		store[user] = AccessRate{Time: time.Now(), Tokens: 3 - 1}
		return "ok - 2"
	} else if accessRate.Tokens > 0 {
		// Decrease tokens.
		store[user] = AccessRate{Time: accessRate.Time, Tokens: accessRate.Tokens - 1}
		return "ok - 3"
	}

	return "sorry"
}
