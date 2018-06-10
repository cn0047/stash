package go_app

import (
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/search"
	"net/http"
	"time"
)

type Visit struct {
	TimeStamp time.Time `datastore:"TimeStamp,noindex"`
	Path      string    `datastore:"Path,noindex"`
}

type SearchUser struct {
	Id        string
	Name      string
	Comment   search.HTML
	Visits    float64
	LastVisit time.Time
	Birthday  time.Time
}

func datastoreHandler(w http.ResponseWriter, r *http.Request) {
	saveVisit(w, r)
	datastorePut1(w, r)
	datastorePut2(w, r)
	transactionCommit(w, r)
	transactionRollBack(w, r)
	transactionPanic(w, r)
	datastoreGetByKey(w, r)
	datastoreGet1(w, r)
	datastoreGet2(w, r)
	indexPut1(w, r)
	indexGet1(w, r)
}

func indexGet1(w http.ResponseWriter, r *http.Request) {
	index, err := search.Open("users")
	if err != nil {
		fmt.Fprintf(w, "<br>Get index 1 fail 1, error: %+v", err)
		return
	}

	ctx := appengine.NewContext(r)
	for t := index.Search(ctx, "<br>SearchUser: Id = user8", nil); ; {
		var doc SearchUser
		id, err := t.Next(&doc)
		if err == search.Done {
			break
		}
		if err != nil {
			fmt.Fprintf(w, "<br>Search error: %v\n", err)
			break
		}
		fmt.Fprintf(w, "<br>Get index 1: %s -> %#v\n", id, doc)
	}
}

func indexPut1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<hr>")

	index, err := search.Open("users")
	if err != nil {
		fmt.Fprintf(w, "<br>Put index 1 fail 1, error: %+v", err)
		return
	}

	u := SearchUser{Id: "usr8", Name: "User 8", Comment: "1 more this is <em>marked up</em> text"}

	ctx := appengine.NewContext(r)
	res, err := index.Put(ctx, "usr8", &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Put index 1 fail 2, error: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>Put index 1 OK: %+v", res)
}

func saveVisit(w http.ResponseWriter, r *http.Request) {
	v := Visit{TimeStamp: time.Now(), Path: "/datastore"}
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Visit", nil)
	key, err := datastore.Put(ctx, key, &v)
	if err != nil {
		fmt.Fprintf(w, "<br>Failed to store visit, error: %+v", err)
	}
	fmt.Fprintf(w, "<br>Visit saved with key 🔑: %+v", key)
}

func createUser(ctx context.Context, id string, name string, tag string) (datastore.Key, User, error) {
	user := User{Id: id, Name: name, Tag: tag}
	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	key := datastore.NewKey(ctx, "User", id, 0, nil)
	key, err := datastore.Put(ctx, key, &user)
	return *key, user, err
}

func datastorePut1(w http.ResponseWriter, r *http.Request) (User, User) {
	u := User{Id: "usr4", Name: "User 4", Tag: "cli", Comment: "this is <em>marked up</em> text"}
	ctx := appengine.NewContext(r)

	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	key := datastore.NewKey(ctx, "User", "user4", 0, nil)
	k, err := datastore.Put(ctx, key, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to PUT: %+v", err)
		return User{}, User{}
	}
	fmt.Fprintf(w, "<br>PUT 1 - OK, key: %+v, 🔑: %+v || %+v", key, k, u)

	u2 := User{}
	err = datastore.Get(ctx, key, &u2)
	if err != nil {
		fmt.Fprintf(w, "<br>Error to GET: %+v", err)
		return User{}, User{}
	}
	fmt.Fprintf(w, "<br>GET * - OK: %+v", u2)

	fmt.Fprintf(w, "<hr>")

	return u, u2
}

func datastorePut2(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key, user, err := createUser(ctx, "x-usr-1", "x1", "x")
	if err != nil {
		fmt.Fprintf(w, "<br>Error to PUT x-usr-1: %+v", err)
		return
	}

	fmt.Fprintf(w, "<br>PUT x-usr-1 - OK: [%+v] %+v", key, user)
}

func transactionCommit(w http.ResponseWriter, r *http.Request) {
	u := User{Id: "usr5", Name: "User 5", Tag: "cli"}
	ctx := appengine.NewContext(r)
	//key := datastore.NewIncompleteKey(ctx, "User", nil)
	key := datastore.NewKey(ctx, "User", "user5", 0, nil)

	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		_, err := datastore.Put(transactionCTX, key, &u)
		// to commit transaction - return nil
		// to rollback transaction - return error
		return err
	}, nil)
	if err == nil {
		fmt.Fprintf(w, "<br>TRANSACTION 1 - ✅ OK, key: %+v | %+v", key, u)
	} else {
		fmt.Fprintf(w, "<br>Transaction 1 🔴 failed, Error: %+v", err)
	}
}

func transactionRollBack(w http.ResponseWriter, r *http.Request) {
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

func transactionPanic(w http.ResponseWriter, r *http.Request) {
	defer func() {
		r := recover()
		fmt.Fprintf(w, "<br>Transaction 3 RECOVERY 🚑: %+v", r)
	}()
	ctx := appengine.NewContext(r)
	err := datastore.RunInTransaction(ctx, func(transactionCTX context.Context) error {
		u := User{Id: "usr7", Name: "User 7", Tag: "cli"}
		key := datastore.NewIncompleteKey(ctx, "User", nil)
		_, err := datastore.Put(transactionCTX, key, &u)
		panic("panic in transaction")
		return err
	}, nil)
	if err != nil {
		fmt.Fprintf(w, "<br>Transaction 3 failed, Error: %+v", err)
	}
}

func datastoreGetByKey(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	k := datastore.NewKey(ctx, "User", "user5", 0, nil)
	u := User{}
	err := datastore.Get(ctx, k, &u)
	if err != nil {
		fmt.Fprintf(w, "<br>Failed get by Key, error: %+v", err)
		return
	}
	fmt.Fprintf(w, "<hr>Get user by key - OK, user: %+v", u)
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
		return
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
