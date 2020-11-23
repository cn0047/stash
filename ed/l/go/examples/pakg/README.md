Info
-

## program

````sh
# install - create hello into `bin` directory:
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c 'cd $GOPATH && go install github.com/cn007b/hello'
# or
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c 'cd $GOPATH/src/github.com/cn007b/hello && go install'

# run
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c '$GOPATH/bin/hello'
````

## library

````sh
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c 'cd $GOPATH && go build github.com/cn007b/stringutil'

docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c 'cd $GOPATH && go install github.com/cn007b/stringutil'
````

## test

````sh
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c 'cd $GOPATH && go test github.com/cn007b/stringutil'
````

## remote packages

````sh
docker run -it --rm -v $PWD:/gh -w /gh -e GOPATH='/gh/ed/go/examples/pkg' \
    golang:latest sh -c 'cd $GOPATH && go get github.com/golang/example/hello'
````
