// Package curl represents tiny wrapper over http and json packages
// and provides opportunity to get decoded JSON response from external REST-API
// service in one line of code.
//
// Example:
//
//	import "github.com/cnkint/curl"
//
//	type GitHubUser struct {
//		Login string `json:"login"`
//	}
//
//	func main() {
//		u := GitHubUser{}
//		err := curl.Unmarshal(curl.Options{URL: "https://api.github.com/users/cn007b"}, &u)
//	}
//
package curl

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// Options contains all possible options.
type Options struct {
	// URL to external REST-API service.
	URL string

	// Timeout specifies a time limit for request to external REST-API service.
	Timeout time.Duration
}

// Unmarshal performs curl request to external REST-API service
// and parses JSON-encoded response
// and stores the result in the value pointed to by p.
func Unmarshal(o Options, p interface{}) error {
	client := http.Client{Timeout: o.Timeout}

	resp, err := client.Get(o.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, p)
	if err != nil {
		return err
	}

	return nil
}
