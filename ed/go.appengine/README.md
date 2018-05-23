App Engine for GOlang
-

[GoLand config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs).

````
"google.golang.org/appengine"
"google.golang.org/appengine/log"
// r *http.Request
log.Infof(appengine.NewContext(r), "%#v", v)

appengine.AppID
appengine.DefaultVersionHostname 
appengine.IsDevAppServer()
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
