App Engine
-

Google Cloud Datastore is a NoSQL document database.

````
gcloud app deploy

gcloud app browse

gcloud app logs tail -s default
````

````
"google.golang.org/appengine"
"google.golang.org/appengine/log"
// r *http.Request
log.Infof(appengine.NewContext(r), "%#v", v)
````

[GoLand config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs).
