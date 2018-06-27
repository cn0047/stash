package go_app

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

type Employee struct {
	Name string
}

type Address struct {
	Address string
}

func indexPutAncestor1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<hr> Ancestor1:")
	ctx := appengine.NewContext(r)

	e := Employee{Name: "Bond"}
	ekey := datastore.NewKey(ctx, "Employee", "bond", 0, nil)
	ek, err := datastore.Put(ctx, ekey, &e)
	if err != nil {
		fmt.Fprintf(w, "<br>indexPutAncestor1 fail 1, error: %+v", err)
		return
	}
	fmt.Fprintf(w, "<br>indexPutAncestor1 put 1 - OK, key: %+v, 🔑: %+v || %+v", ekey, ek, e)

	a := Address{Address: "London"}
	akey := datastore.NewKey(ctx, "Address", "london", 0, ek)
	ak, err := datastore.Put(ctx, akey, &a)
	if err != nil {
		fmt.Fprintf(w, "<br>indexPutAncestor1 fail 2, error: %+v", err)
		return
	}
	fmt.Fprintf(w, "<br>indexPutAncestor1 put 2 - OK, key: %+v, 🔑: %+v || %+v", akey, ak, a)
}

func indexGetAncestor1(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	ekey := datastore.NewKey(ctx, "Employee", "bond", 0, nil)
	q := datastore.NewQuery("Address").Ancestor(ekey)
	a := make([]Address, 0)
	_, err := q.GetAll(ctx, &a)
	if err != nil {
		fmt.Fprintf(w, "<br>indexGetAncestor1 fail 1, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<hr>indexGetAncestor1 - OK: %+v", a)
}
