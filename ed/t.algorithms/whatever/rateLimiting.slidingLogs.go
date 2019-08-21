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
		t := time.Now()
		log[user] = []time.Time{t} // init.
		return fmt.Sprintf("ok - 1 %s", t)
	}

	nt := time.Now()
	log[user] = append(log[user], nt) // add time.

	for len(log[user]) > 0 {
		t := log[user][0]
		diff := time.Since(t)
		if diff.Minutes() >= 1 {
			// delete old log.
			log[user] = log[user][1:]
			fmt.Println("del:", t)
		} else {
			// no use to continue loop, next logs even more fresh.
			break
		}
	}

	if len(log[user]) > 3 {
		return fmt.Sprintf("sorry %s", nt)
	}

	return fmt.Sprintf("ok - 2 %s", nt)
}
