package go_app

import (
	"fmt"
	"github.com/mjibson/goon"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type User struct {
	Id   string `datastore:"-" goon:"id"`
	Name string
	Tag  string
	I    int
	B    bool
	F    float32
	AB   []byte
}

func goonHandler(w http.ResponseWriter, r *http.Request) {
	put1(w, r)
	put2(w, r)
	get1(w, r)
	get3(w, r)
	select1(w, r)
}

func put1(w http.ResponseWriter, r *http.Request) {
	u := &User{Id: "usr1", Name: "Test User1", Tag: "test", I: 7, B: true, F: 3.2, AB: []byte("it works")}
	key, err := goon.NewGoon(r).Put(u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}

	fmt.Fprintf(w, "<br>PUT 1 - OK, key: %+v", key)
}

func put2(w http.ResponseWriter, r *http.Request) {
	u := &User{Id: "usr2", Name: "Test User2", Tag: "test"}
	key, err := goon.NewGoon(r).Put(u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}

	fmt.Fprintf(w, "<br>PUT 2 - OK, key: %+v", key)
}

func get1(w http.ResponseWriter, r *http.Request) {
	u := &User{Id: "usr1"}
	goon.NewGoon(r).Get(u)

	fmt.Fprintf(w, "<br>Get User 1: %+v", u)
}

func get3(w http.ResponseWriter, r *http.Request) {
	u := &User{Id: "usr3"}
	err := goon.NewGoon(r).Get(u)

	fmt.Fprintf(w, "<br>Get User 3: %+v, Error: %+v", u, err)
}

func select1(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := datastore.
		NewQuery("User").
		Filter("Tag =", "test").
		Order("-Name") // order DESC
	u := make([]User, 0)
	_, err := q.GetAll(ctx, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}
	fmt.Fprintf(w, "<br>SELECT 1 - OK: %+v", u)
}
