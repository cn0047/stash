package go_app

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

func datastoreHandler(w http.ResponseWriter, r *http.Request) {
	datastorePut1(w, r)
	datastorePut2(w, r)
	datastorePut3(w, r)
	datastoreGet1(w, r)
	datastoreGet2(w, r)
}

func datastorePut1(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr4", Name: "User 4", Tag: "cli"}
	ctx := appengine.NewContext(r)

	key := datastore.NewIncompleteKey(ctx, "User", nil)
	k, err := datastore.Put(ctx, key, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to PUT: %+v", err)
	}
	fmt.Fprintf(w, "<br>PUT 1 - OK, key: %+v | %+v", k, u)

	u2 := User{}
	err = datastore.Get(ctx, key, &u2)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to GET: %+v", err)
	}
	fmt.Fprintf(w, "<br>GET * - OK, key: %+v | %+v", k, u)

	fmt.Fprintf(w, "<hr>")
}

func datastorePut2(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr5", Name: "User 5", Tag: "cli"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "User", nil)

	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		_, err := datastore.Put(transactionCTX, key, &u)
		// to commit transaction - return nil
		// to rollback transaction - return error
		return err
	}, nil)
	if err == nil {
		fmt.Fprintf(w, "<br>TRANSACTION 1 - OK, key: %+v | %+v", key, u)
	} else {
		fmt.Fprintf(w, "<br>Transaction failed, Error: %+v", err)
	}
}

func datastorePut3(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr6", Name: "User 6 {ROLL BACK CASE}", Tag: "cli"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "User", nil)

	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		datastore.Put(transactionCTX, key, &u)
		err := errors.New("error: test ROLL BACK case")
		return err
	}, nil)
	fmt.Fprintf(w, "<br>TRANSACTION 2: %+v", err)
}

// field tags contains: test & go
func datastoreGet1(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	q := datastore.
		NewQuery("User").
		Filter("Tags =", "test").
		Filter("Tags =", "go")
	u := make([]User, 0)
	_, err := q.GetAll(ctx, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}
	fmt.Fprintf(w, "<hr>SELECT 1 - OK: %+v", u)
}

func datastoreGet2(w http.ResponseWriter, r *http.Request) {

	//entities := make([]User, 0)
	//ctx := appengine.NewContext(r)
	//err = datastore.GetMulti(ctx, []*datastore.Key{"usr1", "usr2"}, entities)
	//q := datastore.GetMulti()
	//
	//u := []User{{Id: "usr1"}, {Id: "usr2"}}
	//ctx := appengine.NewContext(r)
	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	//k, err := datastore.Put(ctx, key, &u)
}
