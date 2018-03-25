GO
-

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
    go get -u github.com/cnkint/curl;
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
````

#### CURL

````
docker run -it --rm -v $PWD:/app -w /app -e GOPATH='/app' golang:latest sh -c '
    go test -v -cover
'
````
