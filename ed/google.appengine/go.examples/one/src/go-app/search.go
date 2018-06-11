package go_app

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/search"
	"net/http"
)

func searchHandler(w http.ResponseWriter, r *http.Request) {
	iput1(w, r)
	iput2(w, r)
	iget1(w, r)
	isearch1(w, r)
}

func iput1(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>iput1 fail 1, error: %+v", err)
		return
	}

	u := SearchUser{Id: "iusr1", Name: "IUser 1", Comment: "One more this is <em>marked up</em> text"}

	ctx := appengine.NewContext(r)
	res, err := index.Put(ctx, "iusr1", &u)
	if err != nil {
		fmt.Fprintf(w, "<br>iput1 fail 2, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>iput1 OK: %+v", res)
}

// iput2 - save in datastore & put into index.
func iput2(w http.ResponseWriter, r *http.Request) {
	id := "iusr2"
	u := SearchUser{Id: id, Name: "IUser 2", Comment: "And this is <em>marked up</em> text"}
	ctx := appengine.NewContext(r)

	key := datastore.NewKey(ctx, "User", id, 0, nil)
	k, err := datastore.Put(ctx, key, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>iput2 error 1: %+v", err)
		return
	}
	fmt.Fprintf(w, "<br>iput1 - OK, key: %+v, 🔑: %+v || %+v", key, k, u)

	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>iput2 fail 2, error: %+v", err)
		return
	}

	res, err := index.Put(ctx, id, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>iput2 fail 3, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>iput2 OK: %+v", res)
}

// iget1 - get by ID.
func iget1(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>iget1 fail 1, error: %+v", err)
		return
	}

	ctx := appengine.NewContext(r)
	id := "iusr2"
	u := SearchUser{}
	err = index.Get(ctx, id, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>iget1 search error: %v\n", err)
		return
	}

	fmt.Fprintf(w, "<br>iget1 OK, result: %s -> %#v\n", id, u)
}

func isearch1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<hr>")
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>isearch1 fail 1, error: %+v", err)
		return
	}

	ctx := appengine.NewContext(r)
	var q string
	//q = "Comment = marked"
	q = "Comment = and"
	for t := index.Search(ctx, q, nil); ; {
		var doc SearchUser
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "<br>isearch1 search error: %v\n", err)
			break
		}
		fmt.Fprintf(w, "<br>isearch1 OK, result: %s -> %#v\n", id, doc)
	}
	fmt.Fprintf(w, "<hr>")
}
