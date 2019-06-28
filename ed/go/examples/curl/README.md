CURL
-

[![Go Report Card](https://goreportcard.com/badge/github.com/plgo/curl)](https://goreportcard.com/report/github.com/plgo/curl)
[![coverage](https://gocover.io/_badge/github.com/plgo/curl?2)](https://gocover.io/github.com/plgo/curl)
[![godoc](https://godoc.org/github.com/plgo/curl?status.svg)](https://godoc.org/github.com/plgo/curl)

`curl` - tiny wrapper over `http` and `json` packages,
which provides opportunity to get decoded JSON response from external REST-API service in one line of code.

## Installation

`go get -u github.com/plgo/curl`

## Usage

````golang
import "github.com/plgo/curl"

type GitHubUser struct {
	Login string `json:"login"`
}

func main() {
	u := GitHubUser{}
	err := curl.Unmarshal(curl.Options{URL: "https://api.github.com/users/cn007b"}, &u)
}
````
