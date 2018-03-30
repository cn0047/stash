App Engine
-

````
export GOPATH=$PWD/ed/go/examples/appEngine

go get -u google.golang.org/appengine/...
go get -u github.com/mjibson/goon

dev_appserver.py --port=8080 --admin_port=8000 --storage_path=$GOPATH/.data --skip_sdk_update_check=true $GOPATH/src/go-app/app.yaml

curl http://localhost:8080
# admin
curl http://localhost:8000

# test
cd src/go-app && /Users/k/.google-cloud-sdk/platform/google_appengine/goroot-1.9/bin/goapp test -cover
````
