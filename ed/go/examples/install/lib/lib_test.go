package lib

import (
    "testing"
)

func TestGetMsg(t *testing.T) {
    e := "It's install."
    a := GetMsg()
    if a != e {
        t.Errorf("Got: %d, want: %d.", a, e)
    }
}
