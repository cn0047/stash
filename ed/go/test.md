Test
-

## Test

````sh
testdata # directory which is ignored by go

# to get more test info
go test -gcflags
go build -gcflags '-m'
go build -gcflags '-m -m' ./main.go

go test -v
go test -cover
go test -race
go test -json
go test -timeout 2s

goapp test -v transactional/users/service -run TestGetServiceAccountsForAdmin

# coverage
go test -coverpkg mylib, fmt mylib
# example
GOPATH=$PWD/ed/go/examples/install
cd $GOPATH/lib ;\
  go test -v ;\
  go test -cover -coverprofile=coverage.out ;\
  go tool cover -html=coverage.out -o=coverage.html ;\
  open coverage.html
````

````
func TestXxx() // test
func BenchmarkXxx() // benchmark
func ExampleXxx() // example (godoc)
````

````golang
import (
  "testing"
)
func TestX(t *testing.T) {
  t.Run("testCase", func(t *testing.T) {
  })
}

t.Log("GivenTheNeedToTestTheSendJSONEndpoint.") {
  //code
}

t.Run("Failed to send email", func(t *testing.T) {})
t.Errorf("Got: %v, want: %v", a, e)
t.Skip("Skipping...")
t.Fail() // mark test as failed
t.FailNow() // mark test as failed and stop execution test further
t.Fatal() // like t.FailNow()
b.RunParallel(func(pb *testing.PB) {})
if testing.Short() { // go test -short }
````

## Bench

````golang
b.ResetTimer() // to reset time and start benchmark (to skip preparations steps)
b.SetParallelism() // sets the number of goroutines used by RunParallel
b.ReportAllocs() // enables malloc statistics
````

Bench output:

````sh
goos: linux
goarch: amd64
BenchmarkFibComplete-4     1000000        2185 ns/op      3911 B/op        0 allocs/op
PASS
ok    _/gh/ed/go/examples/bench 2.446s
````

Means that the loop ran 100000000 times at a speed of 2185 ns (nanosecond) per loop.
