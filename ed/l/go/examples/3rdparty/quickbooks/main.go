package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	BaseURL        = "https://sandbox-quickbooks.api.intuit.com"
	ClientID       = ""
	ConsumerKey    = ""
	ConsumerSecret = ""
	RealmID        = ""
)

var (
	accessToken  = ""
	refreshToken = ""
)

func main() {
	getToken(refreshToken)
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func getToken(refreshToken string) {
	parm := url.Values{}
	parm.Add("grant_type", "refresh_token")
	parm.Add("refresh_token", refreshToken)

	u := "https://oauth.platform.intuit.com/oauth2/v1/tokens/bearer"
	req, err := http.NewRequest(http.MethodPost, u, strings.NewReader(parm.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", getAuthorizationHeader())
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("response Status:", res.Status)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	tokens := Tokens{}
	err = json.Unmarshal(body, &tokens)
	if err != nil {
		panic(err)
	}
	fmt.Printf("tokens: %#v", tokens)

	err = res.Body.Close()
	if err != nil {
		panic(err)
	}
}

func getAuthorizationHeader() string {
	s := ClientID + ":" + ConsumerSecret
	h := base64.StdEncoding.EncodeToString([]byte(s))

	return "Basic " + h
}
