package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	Version   = "1.0"
	Algorithm = "HMAC-SHA1"

	Host = "https://dev.org.com"

	ConsumerKey       = ""
	ConsumerSecretKey = ""
	Token             = ""
	TokenSecret       = ""
)

func main() {
	get(Host+"/api/v2.0/my-groups", nil)
}

func get(urlStr string, reqBody io.Reader) {
	req, err := http.NewRequest(http.MethodGet, urlStr, bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}

	auth := getAuthorizationHeader(http.MethodGet, urlStr)
	req.Header.Set("Authorization", auth)
	fmt.Printf("\n===> %+v \n", urlStr)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("response status:", res.Status)

	respBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("response body:", string(respBody))

	err = res.Body.Close()
	if err != nil {
		panic(err)
	}
}

func getAuthorizationHeader(method, urlStr string) string {
	ts := time.Now().Unix()
	tsStr := fmt.Sprintf("%d", ts)
	nonce := getNonce()
	signature := getSignature(method, urlStr, tsStr, nonce)

	s := `OAuth oauth_consumer_key="` + ConsumerKey + `",`
	s += `oauth_token="` + Token + `",`
	s += `oauth_timestamp="` + tsStr + `",`
	s += `oauth_nonce="` + nonce + `",`
	s += `oauth_version="` + Version + `",`
	s += `oauth_signature_method="` + Algorithm + `",`
	s += `oauth_signature="` + signature + `"`
	//s += `oauth_callback="oob"`
	//s += `application_name="internal"`

	return s
}

func getSignature(method, urlStr, ts, nonce string) string {
	s := method
	//s += url.QueryEscape(urlStr)
	s += urlStr
	s += "?oauth_consumer_key=" + ConsumerKey
	s += "&oauth_token=" + Token
	s += "&oauth_timestamp=" + ts
	s += "&oauth_nonce=" + nonce
	s += "&oauth_version=" + Version
	s += "&oauth_signature_method=" + Algorithm
	//s += "&oauth_callback=oob"

	key := ConsumerSecretKey + "&" + TokenSecret

	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(s))
	res := hex.EncodeToString(h.Sum(nil))

	return res
}

func getNonce() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

	var b strings.Builder
	for i := 0; i < 16; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return b.String()
}
