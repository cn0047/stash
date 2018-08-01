Stackdriver Debugger
-

[doc](https://cloud.google.com/debugger/docs)
[console](https://console.cloud.google.com/debug)

Debugger don't support App Engine.

````bash
wget -O go-cloud-debug https://storage.googleapis.com/cloud-debugger/compute-go/go-cloud-debug
chmod 0755 go-cloud-debug

gcloud debug source gen-repo-info-file --output-directory source

go build -gcflags='-N -l' src/go-app/.gae/main.go
cd src/go-app/.gae && go build -gcflags='-N -l'

export GOPATH=/Users/k/web/kovpak/monitoring
export GOBIN=$GOPATH/bin

./go-cloud-debug \
    -sourcecontext=source/source-context.json \
    -appmodule=default \
    -appversion=20180727t232426 \
    -- go-cloud-debug

/Users/k/.google-cloud-sdk//.install/.backup/platform/google_appengine/goroot-1.9/bin/goapp \
    build src/go-app/.gae/main.go

````

````bash
/Users/k/.google-cloud-sdk//.install/.backup/platform/google_appengine/goroot-1.9/bin/goapp \
    build src/go-app/main.go

````
