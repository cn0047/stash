Google Cloud
-

#### App Engine!

````
cd web/kovpak/gh/ed/go/examples/appEngine

go get -u google.golang.org/appengine/...

dev_appserver.py src/go-app/app.yaml

gcloud app deploy

gcloud app browse

gcloud app logs tail -s default
````
