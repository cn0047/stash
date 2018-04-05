App Engine
-

````
export GOPATH=$PWD/ed/go/examples/appEngine

go get -u google.golang.org/appengine/...
go get -u github.com/mjibson/goon

# test
cd src/go-app && /Users/k/.google-cloud-sdk/platform/google_appengine/goroot-1.9/bin/goapp test -cover

# dev
dev_appserver.py --port=8080 --admin_port=8000 --storage_path=$GOPATH/.data --skip_sdk_update_check=true \
    $GOPATH/src/go-app/app.yaml

# admin
curl http://localhost:8000

# check
curl http://localhost:8080
curl http://localhost:8080/goon

# check config
curl -XGET http://localhost:8080/config



curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"i", "value": 200}'

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"f", "value": 3.14}'

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"s", "value": "100"}'

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"o", "value": {"id": 9, "name": "x"}}'

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"a", "value": ["ok", "true"]}'

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"b", "value": true}'

curl -XPOST http://localhost:8080/config -H 'Content-Type: application/json' \
    -d '{"key":"i", "value": null}'
````
