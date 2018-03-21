/*

docker run -it --rm -v $PWD:/gh -w /gh/ed/go/examples/install \
    golang:latest go run index.go

docker run -it --rm -v $PWD:/gh -w /gh/ed/go/examples/install \
    golang:latest sh -c 'cd lib && go test -v'

docker run -it --rm -v $PWD:/gh -w /gh/ed/go/examples/install \
    golang:latest sh -c 'cd lib && go test -cover -coverprofile=c.out'

# docker run -it --rm -v $PWD:/gh -w /gh/ed/go/examples/install \
#     golang:latest sh -c 'go tool cover -html=lib/c.out -o coverage.html'

*/

package main

import "./lib"

func main() {
    println(lib.GetMsg())
    // println(lib.getCode()) // this won't work
}
