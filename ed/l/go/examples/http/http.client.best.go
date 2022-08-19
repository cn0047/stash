package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	//test()
	real()
}

func real() {
	url := "https://api.github.com/users/cn007b"
	err := one(url, nil)
	if err != nil {
		panic(err)
	}
}

// ok
func test() {
	server := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		_, err := res.Write([]byte(`{"login":"test"}`))
		if err != nil {
		}
	}))
	defer server.Close()

	err := one(server.URL, server.Client())
	fmt.Printf("[test] %+v \n", err)
}

func one(url string, client *http.Client) error {
	params := map[string]string{}
	payload, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("failed to perform JSON marshal, error: %w", err)
	}

	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create new request, error: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	if client == nil {
		client = &http.Client{
			Transport: &http.Transport{
				DialContext:         (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
				TLSHandshakeTimeout: 2 * time.Second,
				MaxIdleConns:        100, // @IMPORTANT: This important for client reusing.
				MaxIdleConnsPerHost: 100, // @IMPORTANT: This important for client reusing.
			},
			Timeout: 5 * time.Second,
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request, error: %w", err)
	}

	// data, err := ioutil.ReadAll(res.Body)
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body, error: %w", err)
	}

	err = res.Body.Close()
	if err != nil {
		return fmt.Errorf("failed to close response body, error: %w", err)
	}

	responseData := make(map[string]interface{})
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return fmt.Errorf("failed to perform JSON unmarshal, error: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("got bad HTTP code: %d", res.StatusCode)
	}

	fmt.Println("response Status:", res.Status)
	fmt.Println("response Headers:", res.Header)
	fmt.Println("response Body length:", len(responseData))
	fmt.Println("login:", responseData["login"])

	return nil
}
