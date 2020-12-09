package main

import (
	"fmt"
	"net/url"
)

func main() {
	simpleQueryString()
}

func simpleQueryString() {
	params := url.Values{}
	params.Add("foo", "bar")
	params.Add("a", "1")
	params.Add("a", "1")
	qs := params.Encode()
	fmt.Printf("Query String: http://.../?%s \n", qs)
}
