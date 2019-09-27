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
	sleep := 100 * time.Millisecond

	fmt.Println(tbrequest("user-1"))
	time.Sleep(sleep)
	fmt.Println(tbrequest("user-1"))
	time.Sleep(sleep)
	fmt.Println(tbrequest("user-1"))
	time.Sleep(sleep)
	fmt.Println(tbrequest("user-1"))
}

// tbrequest allows 3 requests per 1 second.
func tbrequest(user string) string {
	accessRate, ok := store[user]
	if !ok {
		// Init user's tokens.
		store[user] = AccessRate{Time: time.Now(), Tokens: 3 - 1}
		return "ok - 1"
	}

	diff := time.Since(accessRate.Time)
	if diff.Seconds() > 1 {
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
