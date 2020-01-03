trace
-

````golang
package main

import (
  "os"
  "runtime/trace"
)

func main() {
  f, err := os.Create("trace.out")
  if err != nil {
    panic(err)
  }
  defer f.Close()

  err = trace.Start(f)
  if err != nil {
    panic(err)
  }
  defer trace.Stop()

  // Your program here
}
````

````sh
# trace
GOPATH=$PWD/ed/go/examples/install
cd $GOPATH/lib ;\
  go test -v ;\
  go test -trace=trace.out ;\
  go tool trace trace.out
````
