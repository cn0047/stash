package go_app

import (
	"net/http"
	"github.com/mjibson/goon"
	"fmt"
)

type User struct {
	Id    string `datastore:"-" goon:"id"`
	Name  string
}

type Group struct {
	Id   int64 `datastore:"-" goon:"id"`
	Name string
}

func goonHandler(w http.ResponseWriter, r *http.Request) {
	put1(w, r)
	get1(w, r)
	select1(w, r)
}

func put1(w http.ResponseWriter, r *http.Request) {
	u := &User{"usr1", "Test User1"}
	key, err := goon.NewGoon(r).Put(u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}

	fmt.Fprintf(w, "<br>OK, key: %+v", key)
}

func get1(w http.ResponseWriter, r *http.Request) {
	u := &User{Id: "usr1"}
	goon.NewGoon(r).Get(u)

	fmt.Fprintf(w, "<br>User: %+v", u)
}

func select1(w http.ResponseWriter, r *http.Request) {
}
