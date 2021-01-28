package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	f1()
}

func f1() {
	var urls = []string{
		"http://www.1.com/",
		"http://www.2.com/",
		"http://www.3.com/",
	}
	var g errgroup.Group
	for _, url := range urls {
		urlToCheck := url
		g.Go(func() error {
			println(urlToCheck)
			return fmt.Errorf("fail with: %v", urlToCheck)
		})
	}
	err := g.Wait()
	fmt.Printf("err: %#v \n", err) // err: &errors.errorString{s:"fail with: http://www.3.com/"}
}
