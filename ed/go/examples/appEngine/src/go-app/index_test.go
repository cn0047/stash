package go_app

import "testing"

func TestTimeConsuming(t *testing.T) {
	e := "Hello, Gopher Network!"
	a := GeTHomeText()

	if a != e {
		t.Errorf("Got: %s, want: %s", a, e)
	}
}
