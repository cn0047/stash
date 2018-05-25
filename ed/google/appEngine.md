App Engine
-

````
# ~/.google-cloud-sdk/bin/gcloud

# login
gcloud auth login
gcloud auth list

gcloud config list
gcloud config set project thisismonitoring

gcloud projects list

gcloud components list
gcloud components update

gcloud source repos list

gcloud app instances list
gcloud app services list

gcloud app deploy
gcloud app deploy --verbosity=debug --project=thisismonitoring

gcloud app browse

gcloud app logs tail -s default
````

````
# in web console
goapp serve app.yaml
````

#### DataStore

Google Cloud DataStore is a NoSQL document database.
In DataStore nested transactions are not supported.

Disadvantages:

* `FIND ALL WHERE id IN (1, 2)`, `FIND ALL WHERE id = 1 OR id = 2`.
* [How delete element from array](https://monosnap.com/file/YrQHARwcRPAEagaNfoKeMhh1o1bsnZ).

https://cloud.google.com/appengine/docs/standard/go/logs/
