/*

docker run -it --rm -v $PWD:/gh -w /gh/ed/go/examples/install \
    golang:latest go run index.go

*/

package main

import "./lib"

func main() {
    println(lib.GetMsg())
    // println(lib.getCode()) // this won't work
}
