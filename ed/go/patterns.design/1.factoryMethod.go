package main

import (
	"fmt"
)

type Factory struct {
}

func (f Factory) Create(kind string) SocialNetwork {
	var r SocialNetwork

	switch kind {
	case "Twitter":
		r = Twitter{}
	case "Facebook":
		r = Facebook{}
	}

	return r
}

type SocialNetwork interface {
	Share(message string)
}

type Twitter struct {
}

func (t Twitter) Share(message string) {
	fmt.Printf("\nTwitter: %s", message)
}

type Facebook struct {
}

func (f Facebook) Share(message string) {
	fmt.Printf("\nFacebook: %s", message)
}

func main() {
	instance := new(Factory).Create("Twitter")
	instance.Share("rave")
}
