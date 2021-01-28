package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	k         = ""
	s         = ""
	channelId = ""
)

func main() {
	-Get("{id}")
}

func AssignArticleToSection(articleID string, sectionID string) {
}

func GetArticle(articleID string) {
	url := fmt.Sprintf("https://news-api.apple.com/articles/%s", articleID)
	date := time.Now().Format(time.RFC3339)
	canonicalURL := "GET" + url + date

	req, er1 := http.NewRequest("GET", url, nil)
	fmt.Printf("%v\n", er1)
	er2 := addAppleAuthorization(req, canonicalURL, date)
	fmt.Printf("%v\n", er2)

	client := &http.Client{}
	res, er3 := client.Do(req)
	fmt.Printf("%v\n", er3)
	defer res.Body.Close()
	data, er4 := ioutil.ReadAll(res.Body)
	fmt.Printf("%v\n", er4)

	fmt.Printf("%s\n", data)
}

func addAppleAuthorization(r *http.Request, canonicalURL string, date string) error {
	secret, err := encodeSecretKey(canonicalURL, s)
	if err != nil {
		return err
	}

	v := fmt.Sprintf("HHMAC; key=%s; signature=%s; date=%s", k, secret, date)
	r.Header.Set("Authorization", v)

	return nil
}

func encodeSecretKey(message string, secret string) (string, error) {
	key, err := base64.StdEncoding.DecodeString(secret)
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	v := base64.StdEncoding.EncodeToString(h.Sum(nil))

	return v, nil
}
