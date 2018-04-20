package go_app

import (
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
}

func datastorePut1(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr4", Name: "User 4", Tag: "cli"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "User", nil)
	k, err := datastore.Put(ctx, key, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error: %+v", err)
	}

	fmt.Fprintf(w, "<br>PUT 1 - OK, key: %+v | %+v", k, u)
}

func datastorePut2(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr5", Name: "User 5", Tag: "cli"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "User", nil)

	err := datastore.RunInTransaction(ctx, func(ctx context.Context) error {
		_, err := datastore.Put(ctx, key, &u)
		// to commit transaction - return nil
		// to rollback transaction - return not nil
		return err
	}, nil)
	if err == nil {
		fmt.Fprintf(w, "<br>TRANSACTION 1 - OK, key: %+v | %+v", key, u)
		return
	} else {
		fmt.Fprintf(w, "<br>Transaction failed, Error: %+v", err)
		return
	}
}

func datastorePut3(w http.ResponseWriter, r *http.Request) {
}
