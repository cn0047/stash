package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	content = "It works!"
)

func SimpleHTTPHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(content))
}

func Test_SimpleHTTPHandler(t *testing.T) {
	// Create test server.
	ts := httptest.NewServer(http.HandlerFunc(SimpleHTTPHandler))
	defer ts.Close()

	// Make request to test server.
	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	// Asserts.
	if res.StatusCode != http.StatusOK {
		t.Errorf("status code = %d; want %d", res.StatusCode, http.StatusOK)
	}
	actual, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	expected := content
	if string(actual) != expected {
		t.Errorf("response = %q; want %q", actual, expected)
	}
}
