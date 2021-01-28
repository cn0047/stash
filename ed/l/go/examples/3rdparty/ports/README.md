Ports
-

I thought to implement something like [this](https://github.com/cn007b/wall)
or at least like [this](https://github.com/cn007b/monitoring/tree/master/src/go-app)
with purpose to have nice layered architecture,
also thought to use [vo](https://blog.sourcerer.io/go-valueobject-19ea273f9056) throughout
and simplify error handling with [this](https://hackernoon.com/panic-like-a-pro-89044d5a2d35)
but lack of time didn't provide opportunity even add comments...
<br>so, I hope you got the gist what I tried to implement
and we will have broad conversation about all this stuff.

## How to use

To init project run: `make init`.
<br>To test project run: `make test`.
<br>To start project run: `make run`.
<br>Now you can try Rest-API with commands like:
`curl -X GET localhost:8080/ports/ZWUTA`.

To run docker container independently use next commands:

````sh
docker run -it --rm --net=xnet --name srv -p 50051:50051 -v $PWD:/prj -w /prj cn007b/go sh -c '
  export GOPATH=$PWD;
  go run src/ports/app/main.go server;
'

docker run -it --rm --net=xnet -p 8080:8080 -v $PWD:/prj -w /prj cn007b/go sh -c '
  export GOPATH=$PWD;
  go run src/ports/app/main.go client;
'
````

## Dependencies

````sh
google.golang.org/grpc
github.com/golang/protobuf/protoc-gen-go
````
