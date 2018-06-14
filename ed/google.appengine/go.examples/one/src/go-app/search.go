package go_app

import (
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/search"
	"net/http"
	"time"
)

type SearchUser struct {
	Id        string
	Name      string
	Comment   search.HTML
	Visits    float64
	LastVisit time.Time
	Birthday  time.Time
}

// searchHandler - controller for all examples related to search.
func searchHandler(w http.ResponseWriter, r *http.Request) {
	iput1(w, r)
	iput2(w, r)
	iput3(w, r)
	iput4(w, r)

	fmt.Fprintf(w, "<hr>")
	iget1(w, r)

	fmt.Fprintf(w, "<hr><br>🔎 #1: <br>")
	isearch1(w, r)

	fmt.Fprintf(w, "<hr><br>🔎 #2 \"leo\": <br>")
	isearch2(w, r)
}

// iput1 - simple example which put data into index.
func iput1(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>iput1 fail 1, error: %+v", err)
		return
	}

	u := SearchUser{Id: "iusr1", Name: "IUser 1", Comment: "One more time: this is <em>marked up</em> text."}

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
	u := SearchUser{Id: id, Name: "IUser 2", Comment: "And this is <em>marked all up</em> text."}
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

// iput3 - simple example which put into index data related to word "leo".
func iput3(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>iput3 fail 1, error: %+v", err)
		return
	}

	u := SearchUser{Id: "iusr3", Name: "IUser 3",
		Comment: "Theory: Leonardo DiCaprio in Romeo + Juliet Is the Next <b>Big</b> Style Icon."}

	ctx := appengine.NewContext(r)
	res, err := index.Put(ctx, "iusr3", &u)
	if err != nil {
		fmt.Fprintf(w, "<br>iput3 fail 2, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>iput3 OK: %+v", res)
}

// iput4 - simple example which put into index data related to word "leo".
func iput4(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>iput4 fail 1, error: %+v", err)
		return
	}

	u := SearchUser{Id: "iusr4", Name: "IUser 4",
		Comment: "Leo Turns 40 Today! A Look Back at the Star&#039;s <i>Best Roles</i>"}

	ctx := appengine.NewContext(r)
	res, err := index.Put(ctx, "iusr4", &u)
	if err != nil {
		fmt.Fprintf(w, "<br>iput4 fail 2, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>iput3 OK: %+v", res)
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

// isearch1 - simple search example
func isearch1(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>isearch1 fail 1, error: %+v", err)
		return
	}

	ctx := appengine.NewContext(r)
	var q string
	//q = "Comment = marked"
	//q = "Comment = marked up"
	//q = "Comment = \"marked up\""
	//q = "Comment = and"
	//q = "Comment = mark"
	//q = "Comment = arked"
	q = "Comment = ~mark"
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
}

// isearch2 - simple search example which is related to word "leo".
func isearch2(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("user")
	if err != nil {
		fmt.Fprintf(w, "<br>isearch2 fail 1, error: %+v", err)
		return
	}

	ctx := appengine.NewContext(r)
	var q string
	//q = "Comment = ~leo"
	//q = "Comment = leo*"
	q = "Comment = leo"
	for t := index.Search(ctx, q, nil); ; {
		var doc SearchUser
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "<br>isearch2 search error: %v\n", err)
			break
		}
		fmt.Fprintf(w, "<br>isearch2 OK, result: %s -> %#v\n", id, doc)
	}
}
