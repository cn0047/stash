App Engine
-

````
# ~/.google-cloud-sdk/bin/gcloud

gcloud app deploy && gcloud app browse

gcloud app logs tail -s default

gcloud components update
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
