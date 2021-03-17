package go_app

import "testing"

func TestGeTHomeText(t *testing.T) {
	e := "Hello, Gopher Network!"
	a := GeTHomeText()

	if a != e {
		t.Errorf("Got: %s, want: %s", a, e)
	}
}
