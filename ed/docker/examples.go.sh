GO (GOLANG)
-

````
docker build -t xgo ./docker/go

docker run -it --rm -v $PWD:/gh -w /gh xgo go

docker run -it --rm -v $PWD:/gh -w /gh golang:latest go
docker run -it --rm -v $PWD:/gh -w /gh golang:latest go run /gh/x.go

docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh' golang:latest sh -c 'echo $GOPATH'

# test (bench)
docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/go/examples/bench/' \
    golang:latest sh -c 'cd $GOPATH && go test -v'
docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/go/examples/bench/' \
    golang:latest sh -c 'cd $GOPATH && go test -race'
docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/go/examples/bench/' \
    golang:latest sh -c 'cd $GOPATH && go test -v -cpuprofile cpu.out'
# bench
docker run -it --rm -v $PWD:/gh -e GOPATH='/gh/ed/go/examples/bench/' \
    golang:latest sh -c 'cd $GOPATH && go test -bench=. -benchmem'

# db postgresql
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/lib/pq'
# run
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go run src/postgresql/simplest.go'

# db mongo
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go get gopkg.in/mgo.v2'
# or
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go get ./...'
# run
docker run -it --rm --net=xnet -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/db/' \
    golang:latest sh -c 'cd $GOPATH && go run src/mongodb/simple.go'

````

#### Simple Web Server

````
# web.one
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/codegangsta/gin'
# test
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && cd src/firstapp && go test -cover'
docker run -it --rm --name go-one -p 8000:8000 -p 8001:8001 \
    -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.one/' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin --port 8001 --appPort 8000 --path src/firstapp/ run main.go'
# check
curl -i http://localhost:8001/health-check

# web.three ⭐️ ⭐️ ⭐️
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' golang:latest sh -c '
    cd $GOPATH \
    && go get gopkg.in/mgo.v2 \
    && go get github.com/codegangsta/gin \
    && go get -u github.com/derekparker/delve/cmd/dlv
'
# run
docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' \
    golang:latest sh -c 'cd $GOPATH && go run src/app/main.go'
# livereload
docker run -it --rm --net=xnet -p 8080:8080 -p 8081:8081 \
    -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'
# docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/web.three/' \
#     golang:latest sh -c 'cd $GOPATH && ./bin/dlv debug src/app/main.go'

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
# test lr
curl -i -XGET 'http://localhost:8081/cars'
curl -i -XPUT 'http://localhost:8081/cars'
curl -i -XDELETE 'http://localhost:8081/cars/1'
curl -i -XPOST 'http://localhost:8081/cars' -H 'Content-Type: application/json' \
   -d '{"vendor": "BMW", "name": "M6"}'
````

#### GO Echo

````
# one

# init
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' golang:latest sh -c '
    cd $GOPATH \
    && go get -u github.com/labstack/echo/... \
    && go get -u github.com/codegangsta/gin
'
# run
docker run -it --rm -p 8080:8080 -p 8081:8081 \
    -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' \
    golang:latest sh -c 'cd $GOPATH && ./bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'
# test
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.echo/examples/one' golang:latest sh -c '
    cd $GOPATH && cd src/app && go test -cover
'
# check
curl -i -XGET 'http://localhost:8081'
curl -i -XGET 'http://localhost:8081/products'
curl -i -XGET 'http://localhost:8081/products/iphone'
````

#### GO Gin

````
docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/gin-gonic/gin'

docker run -it --rm -p 8080:8080 -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go.gin/examples/one' \
    golang:latest sh -c 'cd $GOPATH && go run src/one/main.go'

# curl localhost:8080/v1/file-info/id/7
````

#### API-Gateway

````
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go get github.com/codegangsta/gin;
    go get -u golang.org/x/lint/golint;
'
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go get -v -t -d ./...
'
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go fmt ./...
'
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/app/ && go vet ./...
'
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    ./bin/golint src/app/...
'
# run
docker run -it --rm -p 8080:8080 -p 8081:8081 \
    -v $PWD:/app -w /app -e GOPATH='/app' \
    golang:latest sh -c './bin/gin --port 8081 --appPort 8080 --path src/app/ run main.go'

curl 'http://localhost:8081/github/users/cn007b'|jq

# install
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/app/ && go install
'
# run
docker run -it --rm -p 8080:8080  -v $PWD/bin/app:/app/app -w /app cn007b/ubuntu ./app
# check
curl 'http://localhost:8081/github/users/cn007b'|jq
````

#### CURL

````
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go test -v -cover
'
````

#### PRA

````
# 1) Install dependencies:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go get -u golang.org/x/lint/golint;
    go get -u golang.org/x/tools/cmd/cover;
    go get github.com/mattn/goveralls
'

# 2) Run tests:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/app/products/ && go test -v ./...
'

# 3) Run format:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/app/products/ && go fmt ./...
'

# 4) Run vet:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/app/products/ && go vet ./...
'

# 5) Run lint:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    ./bin/golint src/github.com/app/products/...
'

# 6) Check manually:
docker run -it --rm -p 8080:8080 -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go run src/github.com/app/products/main.go
'
# Check:
curl -XGET 'http://localhost:8080'
curl -XGET 'http://localhost:8080/products' | jq
curl -XGET 'http://localhost:8080/products/mTd3lb' | jq
````

````
# Coverage one file:
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/app/products/dao && go test -cover -coverprofile=coverage.out ./...
'
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    cd src/github.com/app/products/dao && go tool cover -html=coverage.out -o=coverage.html
'
````
