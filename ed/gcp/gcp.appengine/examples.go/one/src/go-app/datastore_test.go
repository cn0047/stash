package go_app

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine/aetest"
	"testing"
)

func GetTestContext(t *testing.T) context.Context {
	t.Helper()

	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	return ctx
}

func TestCreateUser(t *testing.T) {
	_, _, err := createUser(GetTestContext(t), "t-user-1", "TUser1", "test")
	if err != nil {
		t.Errorf("Got: %v, want: %v", err, nil)
	}
}
