# go

docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh' xgo /bin/bash

# debug
export APP_DIR='/gh/ed/l/go/examples/debug'
export GOPATH=$PWD/..$APP_DIR
docker run -it --rm -v $PWD:/gh -e GOPATH=$APP_DIR xgo sh -c '
  cd $GOPATH \
  && go get -u github.com/derekparker/delve/cmd/dlv
'
# run
docker run -it --rm -p 8080:8080 -v $PWD:/gh -e GOPATH=$APP_DIR xgo sh -c '
  cd $GOPATH && go run src/app/main.go
'

# whatever - slice
export GOPATH=$PWD/ed/l/go/examples/whatever/_slice.allocation
docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/l/go/examples/whatever/_slice.allocation/' xgo sh -c '
  cd $GOPATH \
  && go get -u github.com/google/pprof \
  && go get -u github.com/pkg/profile
'
# run bench
docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/l/go/examples/whatever/_slice.allocation/' \
  xgo sh -c 'cd $GOPATH/src/app/lib && go test -bench=. -benchmem'
# # install
# docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/l/go/examples/whatever/_slice.allocation/' \
#   xgo sh -c 'cd $GOPATH/src/app/ && go install'
# # run
# docker run -it --rm -p 8000:8000 -v $PWD:/gh -v $PWD/ed/l/go/examples/whatever/_slice.allocation/tmp:/tmp \
#   -e GOPATH='/gh/ed/l/go/examples/whatever/_slice.allocation/' \
#   xgo sh -c 'cd $GOPATH/bin && ./app'
# or
docker run -it --rm -p 8000:8000 -v $PWD:/gh -e GOPATH='/gh/ed/l/go/examples/whatever/_slice.allocation/' \
  xgo sh -c '
    cd $GOPATH/src/app/ \
    && go build \
    && ./app
  '
cd $GOPATH/src/app/
go tool pprof app http://localhost:8000/debug/pprof/profile
# run apache bench
sudo ifconfig lo0 alias 10.254.254.254
docker run -ti --rm xubuntu ab -k -c 8 -n 100000 "http://10.254.254.254:8000/f1"
go tool pprof app http://localhost:8000/debug/pprof/heap
go tool pprof app http://localhost:8000/debug/pprof/goroutine
go tool pprof app http://localhost:8000/debug/pprof/block
# check
curl http://localhost:8000/f1
curl http://localhost:8000/f2
open http://localhost:8000/debug/pprof
# cli:
# For cli have to generate cpu.out report from unit test
# and run pprof command with generated cpu.out report.

# test (bench)
export APP_PATH=ed/go/examples/bench/fibonacci/
export GOPATH=/gh/$APP_PATH
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH/.. && go run main.go'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -v'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -v  -parallel 9'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -v -cpu=1'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -v -cpu=2'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -cover'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -bench=. -benchmem -benchtime 1m1s'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -bench=. -benchmem -cpu=2'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c 'cd $GOPATH && go test -race'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c '
  cd $GOPATH \
  && go test -v -cpuprofile cpu.out \
  && go tool pprof -png -output report.cpu.png fibonacci.test cpu.out
'
# pprof -http=0.0.0.0:8080
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c '
  cd $GOPATH \
  && go test -v -memprofile mem.out \
  && go tool pprof -png -output report.mem.png fibonacci.test mem.out
'
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c '
  cd $GOPATH \
  && go test -v -mutexprofile mtx.out \
  && go tool pprof -png -output report.mtx.png fibonacci.test mtx.out
'
# or
export GOPATH=$PWD/ed/l/go/examples/bench/fibonacci
cd $GOPATH
go test -v -run TestFibSimple
go test -bench=. -benchmem -cpu=2
go test -v -memprofile mem.out -run TestFibSimple
go tool pprof -png -output report.mem.png fibonacci.test mem.out

# blur
go run ed/l/go/examples/blur/app/bench.go
# or
# docker run -it --rm -v $PWD:/gh -w /gh/ed/l/go/examples/blur -e GOPATH=$GOPATH xgo sh -c '
#   go run app/bench.go
# '

# db postgresql
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go get github.com/lib/pq'
# run
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go run src/postgresql/simplest.go'
# local
GOPATH=$PWD/ed/l/go/examples/db
go run $GOPATH/src/postgresql/simplest.go

# db mongo
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go get gopkg.in/mgo.v2'
# or
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go get ./...'
# run
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go run src/mongodb/simple.go'

# db mysql
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go get github.com/go-sql-driver/mysql'
# run
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go run src/mysql/simple.go'

# mysql-wear
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' \
  xgo sh -c 'cd $GOPATH && go run src/mysql-wear/simple.go'

# neo4j
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/db/' cn007b/go:1.13-neo4j sh -c '
  cd $GOPATH; \
  go get github.com/neo4j/neo4j-go-driver/neo4j; \
  go run src/neo4j/simple.go
'

# grpc one
docker run -it --rm --name=grpcone -v $PWD:/gh -w /gh/ed/l/go/examples/grpc/one xgo sh -c '
  export GOPATH=$PWD;
  go get -u google.golang.org/grpc;
  go get -u github.com/golang/protobuf/protoc-gen-go;
  export PATH=$PATH:$GOPATH/bin;
  protoc -I src/app/lib src/app/lib/one.proto --go_out=plugins=grpc:src/app/lib;
  go run src/app/server.go;
'
# &
docker exec -it grpcone sh -c 'export GOPATH=$PWD; go run src/app/client.go'

# youtube
GOPATH=$PWD/ed/l/go/examples/youtube


# redis
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -e GOPATH=/gh/ed/l/go/examples/redis/app xgo sh -c '
  cd $GOPATH && go run main.go'
# or
GOPATH=$PWD/ed/l/go/examples/redis
cd $GOPATH
go run app/main.go

# web.one
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/web.one/' xgo sh -c '
  cd $GOPATH \
  && go get github.com/codegangsta/gin \
  && go get github.com/pkg/errors
'
# test
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/web.one/' \
  xgo sh -c 'cd $GOPATH && cd src/firstapp && go test -cover'
# run
docker run -it --rm --name go-one -p 8000:8000 -p 8001:8001 \
  -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/web.one/' \
  xgo sh -c 'cd $GOPATH && ./bin/gin --port 8001 --appPort 8000 --path src/firstapp/ run main.go'
# check
curl -i http://localhost:8001/health-check

# web.three ⭐️⭐️⭐️
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/web.three/' xgo sh -c '
  cd $GOPATH \
  && go get gopkg.in/mgo.v2 \
  && go get github.com/codegangsta/gin \
  && go get -u github.com/derekparker/delve/cmd/dlv
'
# run
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/web.three/' \
  xgo sh -c 'cd $GOPATH && go run src/app/main.go'
# livereload
docker run -it --rm --net=xnet -p 8080:8080 -p 8081:8081 \
  -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/l/go/examples/web.three/' \
  xgo sh -c 'cd $GOPATH && ./bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'
# test
curl -i 'http://localhost:8080'
curl -i 'http://localhost:8080/home'
curl -i 'http://localhost:8080/cars'
curl -i -XGET 'http://localhost:8080/cars'
curl -i -XPUT 'http://localhost:8080/cars'
curl -i -XDELETE 'http://localhost:8080/cars'
curl -i -XPOST 'http://localhost:8080/cars'
curl -i -XPOST 'http://localhost:8080/cars' -H 'Content-Type: application/json' \
   -d '{"vendor": "BMW", "name": "X5"}'
# test live-reload
curl -i -XGET 'http://localhost:8081/cars'
curl -i -XPUT 'http://localhost:8081/cars'
curl -i -XDELETE 'http://localhost:8081/cars/1'
curl -i -XPOST 'http://localhost:8081/cars' -H 'Content-Type: application/json' \
   -d '{"vendor": "BMW", "name": "M6"}'

# web.three.tiny
docker run -it --rm -v $PWD/ed/l/go/examples/web.three.tiny:/app -e GOPATH='/app' \
  cn007b/go sh -c 'cd $GOPATH/src/app && go install'
docker run -it --rm -v $PWD/ed/l/go/examples/web.three.tiny:/app -e GOPATH='/app' \
  -p 8080:8080 \
  cn007b/go sh -c 'cd $GOPATH && ./bin/app'
# or
GOPATH=$PWD/ed/l/go/examples/web.three.tiny
cd $GOPATH
go run src/app/main.go
go install src/app
go build src/app
# check
curl http://localhost:8080/v1/id/7

# algolia
export GOPATH=$PWD'/ed/l/go/examples/algolia'
go get github.com/algolia/algoliasearch-client-go/algoliasearch
go run ed/go/examples/algolia/src/main.go

# algolia
export GOPATH=$PWD'/ed/l/go/examples/airbrake'
go get github.com/airbrake/gobrake
go run ed/go/examples/airbrake/src/app/main.go

# aws
# Init AWS with terraform first,
# @see: ed/sh/sh.terraform/README.md
export GOPATH=$PWD/ed/l/go/examples/aws
cd $GOPATH/src/app
go get -u "github.com/aws/aws-sdk-go/aws"
go get -u "github.com/aws/aws-sdk-go/aws/arn"
go get -u "github.com/stretchr/testify/assert"
go get -u "github.com/thepkg/awsl"
# aws lambda
# @see: ed/sh/terraform/README.md
go test -v ./lambda
# dynamodb
go run main.go k21 v21 21
go run main.go k5
# aws s3
# aws sns
# aws sqs
go run main.go

# gcp
cd $PWD/ed/l/go/examples/3rdparty/gcp
go run main.go

# jwt
export GOPATH=$PWD/ed/l/go/examples/jwt
cd $GOPATH
go get -u "github.com/dgrijalva/jwt-go"
go run src/app/main.go

# bench
GOPATH=$PWD/ed/l/go/examples/bench
go run $GOPATH/main.go

# websocket
export GOPATH=$PWD'/ed/l/go/examples/websocket'
go get -u github.com/gorilla/websocket
cd $GOPATH/src/app && go run $GOPATH/src/app/websockets.go

# image
i=./ed/l/php/php.yii/examples/testdrive/css/bg.gif
i=./ed/l/go/examples/3rdparty/aws/src/app/s3/s.png
i=./ed/l/nodejs/examples/mongo.university/mongomart/static/img/logo.jpg
i=./ed/security/exploit.seh.jpeg
go run ed/l/go/examples/image/main.go $i

# debug
export GOPATH=$PWD'/ed/l/go/examples/debug'
# install delve
cd $GOPATH && git clone https://github.com/derekparker/delve.git && cd -
cd $GOPATH/delve && make install && cd -
#
go build -gcflags='-N -l' $GOPATH/src/app/main.go \
  && dlv --listen=:2345 --headless=true --api-version=2 exec ./main
#
# https://monosnap.com/file/xQWOFnKzTy2ODUu4Kxute5nxSEtTuR
#
# ⬇️
# docker run -it --rm -p 2345:2345 -v $GOPATH:/gh -w /gh -e GOPATH='/gh' xgo sh -c '
#   go build -gcflags="-N -l" src/app/main.go \
#   && /app/bin/dlv --listen=:2345 --headless=true --api-version=2 exec ./main
# '
#
# ERROR:
# could not launch process: fork/exec ./main: operation not permitted



#### GO Echo

# one

# init
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' xgo sh -c '
  cd $GOPATH \
  && go get -u github.com/labstack/echo/... \
  && go get -u github.com/codegangsta/gin
'
# run
docker run -it --rm -p 8080:8080 -p 8081:8081 \
  -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' \
  xgo sh -c 'cd $GOPATH && ./bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'
# test
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' xgo sh -c '
  cd $GOPATH && cd src/app && go test -cover
'
# check
curl -i -XGET 'http://localhost:8081'
curl -i -XGET 'http://localhost:8081/products'
curl -i -XGET 'http://localhost:8081/products/iphone'

#### GO Gin

docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
  xgo sh -c 'cd $GOPATH && go get github.com/gin-gonic/gin'
# or
docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
  xgo sh -c 'cd $GOPATH && go get ./...'

docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
  xgo sh -c 'cd $GOPATH && go run src/one/main.go'
# or inside container with /bin/bash
# go get ./...
# go install ./...

# curl localhost:8080/v1/id/7

#### API-Gateway

docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  go get github.com/codegangsta/gin;
  go get -u golang.org/x/lint/golint;
  go get -u github.com/thepkg/rest;
'
# docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
#   go get -v -t -d ./...
# '
# docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
#   go fmt ./...
# '
# docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
#   cd src/app/ && go vet ./...
# '
# docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
#   ./bin/golint src/app/...
# '
# run
docker run -it --rm -p 8080:8080 -p 8081:8081 \
  -v $PWD:/app -w /app -e GOPATH='/app' \
  xgo sh -c './bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'

curl 'http://localhost:8081/github/users/cn007b'|jq

# install
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  cd src/app/ && go install
'
# run
docker run -it --rm -p 8080:8080  -v $PWD/bin/app:/app/app -w /app cn007b/ubuntu ./app
# check
curl 'http://localhost:8081/github/users/cn007b'|jq

#### PRA

# 1) Install dependencies:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  go get -u golang.org/x/lint/golint;
  go get -u golang.org/x/tools/cmd/cover;
  go get github.com/mattn/goveralls
'
# 2) Run tests:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  cd src/github.com/app/products/ && go test -v ./...
'
# 3) Run format:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  cd src/github.com/app/products/ && go fmt ./...
'
# 4) Run vet:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  cd src/github.com/app/products/ && go vet ./...
'
# 5) Run lint:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  ./bin/golint src/github.com/app/products/...
'
# 6) Check manually:
docker run -it --rm -p 8080:8080 -v $PWD:/app -w /app -e GOPATH='/app' xgo sh -c '
  go run src/github.com/app/products/main.go
'
# Check:
curl -XGET 'http://localhost:8080'
curl -XGET 'http://localhost:8080/products' | jq
curl -XGET 'http://localhost:8080/products/mTd3lb' | jq

#### Products (PRA)

export APP_PATH=$PWD/ed/go.echo/examples/pra

docker run -it --rm -v $APP_PATH:/app -e GOPATH='/app' xgo sh -c '
  cd $GOPATH/src/github.com/app/products/dao \
  && go get -u github.com/labstack/echo/...
'

# Coverage one file:
docker run -it --rm -v $APP_PATH:/app -e GOPATH='/app' xgo sh -c '
  cd $GOPATH/src/github.com/app/products/dao \
  && go test -cover -coverprofile=coverage.out ./... \
  && go tool cover -html=coverage.out -o=coverage.html
'
docker run -it --rm -v $APP_PATH:/app -e GOPATH='/app' xgo sh -c '
  cd $GOPATH/src/github.com/app/products \
  && mkdir -p ./.cover \
  && go test -cover -coverprofile ./.cover/a.part ./config \
  && go test -cover -coverprofile ./.cover/b.part ./controller \
  && go test -cover -coverprofile ./.cover/c.part ./dao \
  && go test -cover -coverprofile ./.cover/d.part ./di \
  && echo "mode: set" > coverage.out \
  && grep -h -v "mode: set" .cover/*.part >> coverage.out \
  && go tool cover -html=coverage.out -o=coverage.html
'
docker run -it --rm -v $APP_PATH:/app -e GOPATH='/app' xgo sh -c '
  cd $GOPATH/src/github.com/app/products \
  && mkdir -p ./.cover \
  && go test -cover -coverprofile ./.cover/config.part ./config \
  && go test -cover -coverprofile ./.cover/controller.part ./controller \
  && go test -cover -coverprofile ./.cover/dao.part ./dao \
  && go test -cover -coverprofile ./.cover/di.part ./di \
  && echo "mode: set" > ./.cover/coverage.out \
  && grep -h -v "mode: set" .cover/*.part >> ./.cover/coverage.out \
  && go tool cover -html=./.cover/coverage.out -o=./.cover/coverage.html
'

#### (thepkg) strings & gcd & rest

docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' cn007b/go:1.10 go get -v -t -d ./...
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' cn007b/go:1.10 go vet
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' cn007b/go:1.10 go fmt ./...
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' cn007b/go:1.10 golint ./...
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' cn007b/go:2018-06-07 sh -c '
  go test -v -covermode=count -coverprofile=coverage.out
'
# docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' cn007b/go:1.10 gometalinter ./...

#### GO AppEngine all

export GOPATH=$HOME/web/kovpak/gh/ed/cloud/gcp/gcp.appengine/examples.all/go
cd $GOPATH
$HOME/.gcloud/bin/gcloud app deploy -q src/go-app/app.yaml

#### GO AppEngine one

export GOPATH=$PWD/ed/google.appengine/go.examples/one
cd $GOPATH

go get -u google.golang.org/appengine/...
go get -u github.com/mjibson/goon
go get -u golang.org/x/lint/golint
go get -u golang.org/x/tools/cmd/cover
go get -u github.com/google/gops
go get -u github.com/kisielk/godepgraph
go get -u github.com/thepkg/strings
go get -u github.com/thepkg/gcd

# test
cd $GOPATH/src/go-app && ~/.google-cloud-sdk/platform/google_appengine/goroot-1.9/bin/goapp test -cover

# check
curl http://localhost:8080/goon

# godepgraph
cd $GOPATH && godepgraph google.golang.org/appengine | dot -Tpng -o godepgraph.png
cd $GOPATH && godepgraph github.com/thepkg/strings | dot -Tpng -o godepgraph.png

# unit test
cd $GOPATH/src/go-app && goapp test ./...

export GOPATH=/gh/ed/google.appengine/go.examples/one
docker run -it --rm -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c '
  cd $GOPATH && go get ./...
'
docker run -it --rm -p 8000:8000 -p 8001:8001 -v $PWD:/gh -e GOPATH=$GOPATH xgo sh -c '
  dev_appserver.py --log_level=debug --host=0.0.0.0 --port=8000 --admin_host=0.0.0.0 --admin_port=8001 \
  --storage_path=$GOPATH/.data --skip_sdk_update_check=true --support_datastore_emulator=no \
  $GOPATH/src/go-app/app.yaml
'

#### Monitoring

# export GOROOT=/Users/k/.google-cloud-sdk/platform/google_appengine/goroot-1.9
# export GOPATH=/Users/k/web/kovpak/monitoring:/Users/k/web/kovpak/monitoring/src/go-app

go get ./src/go-app/...
go get -u github.com/thepkg/strings

# for circleci
docker run -it --rm -v $PWD:/app -w /app -e GOPATH=/app cn007b/go sh -c '
  cd $GOPATH/src/go-app && go vet
'
docker run -it --rm -v $PWD:/app -w /app cn007b/go golint src/go-app/...
docker run -it --rm -v $PWD:/app -w /app -e GOPATH=/app cn007b/go sh -c '
  cd $GOPATH/src/go-app && go fmt ./...
'

# deploy PROD
gcloud auth login
gcloud config set project itismonitoring
gcloud config set disable_usage_reporting false
gcloud config list
gcloud app versions list
for i in $( gcloud app versions list | grep STOPPED | awk '{print $2}' ); do
  gcloud app versions delete -q $i
done
export GOPATH=/Users/k/web/kovpak/monitoring
gcloud app deploy -q src/go-app/.gae/app.yaml
gcloud app deploy -q src/go-app/.gae/cron.yaml
gcloud app deploy -q src/go-app/.gae/queue.yaml
gcloud app deploy -q src/go-app/.gae/index.yaml
