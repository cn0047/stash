App Engine for GOlang
-

````
"google.golang.org/appengine"
"google.golang.org/appengine/log"

ctx := appengine.NewContext(r)
module := appengine.ModuleName(ctx)
instance := appengine.InstanceID()

// r *http.Request
log.Infof(appengine.NewContext(r), "%#v", v)

appengine.AppID
appengine.DefaultVersionHostname 
appengine.IsDevAppServer()
````

#### DataStore

comparison operators: `=, <, <=, >, >=`.

````
key := datastore.NewKey(
        ctx,        // context.Context
        "Employee", // Kind
        "asalieri", // String ID; empty means no string ID
        0,          // Integer ID; if 0, generate automatically. Ignored if string ID specified.
        nil,        // Parent Key; nil means no parent
)

_, err = datastore.PutMulti(ctx, []*datastore.Key{k1, k2, k3}, []interface{}{e1, e2, e3})

var entities = make([]*T, 3)
err = datastore.GetMulti(ctx, []*datastore.Key{k1, k2, k3}, entities)

err = datastore.DeleteMulti(ctx, []*datastore.Key{k1, k2, k3})

q := datastore.NewQuery("Person").KeysOnly()
````

#### Local Unit Testing

````
# run
goapp test ./demos/transaction

"google.golang.org/appengine/aetest"

ctx, done, err := aetest.NewContext()

inst, err := aetest.NewInstance(nil)
defer inst.Close()
````
