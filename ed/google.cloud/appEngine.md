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
