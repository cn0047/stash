App Engine
-

[GoLand config](https://monosnap.com/file/X5w1jrpQ1C4fSmn7rmU9Lbm0l3xNBs).

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

````
export GOPATH=$PWD/ed/go/examples/appEngine
cd $GOPATH

go get -u google.golang.org/appengine/...
go get -u github.com/mjibson/goon
go get -u golang.org/x/lint/golint
go get -u golang.org/x/tools/cmd/cover

# test
cd src/go-app && /Users/k/.google-cloud-sdk/platform/google_appengine/goroot-1.9/bin/goapp test -cover

# start dev server
dev_appserver.py \
    --port=8080 --admin_port=8000 --storage_path=$GOPATH/.data --skip_sdk_update_check=true \
    $GOPATH/src/go-app/app.yaml
# or
~/.google-cloud-sdk/bin/dev_appserver.py \
    --port=8080 --admin_port=8000 --storage_path=$GOPATH/.data --skip_sdk_update_check=true \
    $GOPATH/src/go-app/app.yaml

# admin
curl http://localhost:8000

# check
curl http://localhost:8080
curl http://localhost:8080/goon
````

````
# check configs:

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"i", "value": 200, "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"f", "value": 3.14, "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"s", "value": "100", "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"o", "value": {"id": 9, "name": "x"}}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"o2", "value": {"id": "config", "value": {"n": 1, "b": false, "s": "ok"}}}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"a", "value": ["ok", "true"]}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"b", "value": true, "tag": "simple"}' | jq

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"n", "value": null, "tag": "simple"}' | jq

curl -XGET http://localhost:8080/config/i | jq
curl -XGET http://localhost:8080/config/f | jq
curl -XGET http://localhost:8080/config/s | jq
curl -XGET http://localhost:8080/config/o | jq
curl -XGET http://localhost:8080/config/o2 | jq
curl -XGET http://localhost:8080/config/a | jq
curl -XGET http://localhost:8080/config/b | jq
curl -XGET http://localhost:8080/config/n | jq

curl -XGET http://localhost:8080/config/tag/simple | jq

// TODO:
// 1) curl -XGET http://localhost:8080/config/?q=tag1:name;tag2:*;tag3:footer,header;tagPrefix*:copyright
// 2) cache response payload
````
