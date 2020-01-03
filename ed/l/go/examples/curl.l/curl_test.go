package curl

import (
	"testing"
)

// GitHubUser - Structure which represents github user profile.
type GitHubUser struct {
	Login string `json:"login"`
}

// TestBadURL test case for bad URL.
func TestBadURL(t *testing.T) {
	expected := `Get error://error: unsupported protocol scheme "error"`

	u := GitHubUser{}
	err := Unmarshal(Options{URL: "error://error"}, &u)
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestNotJSON test case for not JSON response.
func TestNotJSON(t *testing.T) {
	expected := `invalid character '<' looking for beginning of value`

	u := GitHubUser{}
	err := Unmarshal(Options{URL: "https://github.com"}, &u)
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestNotFound test case for not found user.
func TestNotFound(t *testing.T) {
	expected := `404 Not Found`

	u := GitHubUser{}
	err := Unmarshal(Options{URL: "https://api.github.com/users/{error}"}, &u)
	actual := err.Error()

	if actual != expected {
		t.Errorf("Got: %s, want: %s.", actual, expected)
	}
}

// TestJSON test case for success JSON Unmarshal.
func TestJSON(t *testing.T) {
	u := GitHubUser{}
	err := Unmarshal(Options{URL: "https://api.github.com/users/cn007b"}, &u)

	if err != nil {
		t.Errorf("Got: %s, want: nil.", err.Error())
	}
	if u.Login != "cn007b" {
		t.Errorf("Got: %s, want: cn007b.", u.Login)
	}
}
