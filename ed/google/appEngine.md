App Engine
-

[logs](https://console.cloud.google.com/logs)

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
gcloud components install app-engine-php

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

Warmup Requests - you can use to avoid latency while loading application code on a fresh instance.

## DataStore

Google Cloud DataStore is a NoSQL document database.
In DataStore nested transactions are not supported.

Disadvantages:

* `FIND ALL WHERE id IN (1, 2)`, `FIND ALL WHERE id = 1 OR id = 2`.
* [How delete element from array](https://monosnap.com/file/YrQHARwcRPAEagaNfoKeMhh1o1bsnZ).

## Examples

#### GO one

````
export GOPATH=$PWD/ed/go.appengine/examples/one
# cd $GOPATH

go get -u google.golang.org/appengine/...
go get -u github.com/mjibson/goon
go get -u golang.org/x/lint/golint
go get -u golang.org/x/tools/cmd/cover
go get -u github.com/google/gops
go get -u github.com/kisielk/godepgraph
go get -u github.com/thepkg/strings

# test
cd $GOPATH/src/go-app && ~/.google-cloud-sdk/platform/google_appengine/goroot-1.9/bin/goapp test -cover

# start dev server
~/.google-cloud-sdk/bin/dev_appserver.py \
    --port=8080 --admin_port=8000 --storage_path=$GOPATH/.data --skip_sdk_update_check=true \
    $GOPATH/src/go-app/app.yaml

# check
curl http://localhost:8080/goon

# godepgraph
cd $GOPATH && godepgraph google.golang.org/appengine | dot -Tpng -o godepgraph.png
cd $GOPATH && godepgraph github.com/thepkg/strings | dot -Tpng -o godepgraph.png

# unit test
cd $GOPATH/src/go-app && goapp test ./...
````

#### Monitoring

````
# export GOROOT=/Users/k/.google-cloud-sdk/platform/google_appengine/goroot-1.9
export GOPATH=/Users/k/web/kovpak/monitoring
# export GOPATH=/Users/k/web/kovpak/monitoring:/Users/k/web/kovpak/monitoring/src/go-app

go get ./src/go-app/...

~/.google-cloud-sdk/bin/dev_appserver.py \
    --port=8080 --admin_port=8000 --storage_path=$GOPATH/.data --skip_sdk_update_check=true \
    $GOPATH/src/go-app/app.yaml

# deploy PROD
gcloud config set project thisismonitoring
cd src/go-app && gcloud app deploy
````

#### PHP one

````
cd /Users/k/web/kovpak/gh/ed/php.appengine/examples/one
composer install

~/.google-cloud-sdk/bin/dev_appserver.py \
    --port=8080 --admin_port=8000 --skip_sdk_update_check=true app.yaml

gcloud config set project thisissimplebot
gcloud app deploy
````
