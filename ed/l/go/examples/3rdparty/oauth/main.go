package main

import (
	"fmt"
	"io/ioutil"

	"github.com/dghubble/oauth1"
)

const (
	APIKey         = ""
	APISecretKey   = ""
	APIToken       = ""
	APITokenSecret = ""
)

func main() {
	dghubbleSimple()
}

func dghubbleSimple() {
	config := oauth1.NewConfig(APIKey, APISecretKey)
	config.Signer = &oauth1.HMACSigner{ConsumerSecret: APISecretKey}
	config.Signer = &oauth1.HMAC256Signer{ConsumerSecret: APISecretKey}
	token := oauth1.NewToken(APIToken, APITokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	url := "https://api.twitter.com/1.1/statuses/home_timeline.json?count=2"
	resp, err := httpClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Raw Response Body:\n%v\n\n", string(body))
	fmt.Printf("Raw Response:\n%v\n", resp)
}
