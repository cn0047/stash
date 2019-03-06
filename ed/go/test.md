Test
-

## Test

````sh
# to get more test info
go test -gcflags
go build -gcflags '-m'
go build -gcflags '-m -m' ./main.go

go test -v
go test -cover
go test -race
go test -json

goapp test -v transactional/users/service -run TestGetServiceAccountsForAdmin

# coverage
GOPATH=$PWD/ed/go/examples/install
cd $GOPATH/lib ;\
  go test -v ;\
  go test -cover -coverprofile=coverage.out ;\
  go tool cover -html=coverage.out -o=coverage.html ;\
  open coverage.html
````

````golang
import (
  "testing"
)
func TestX(t *testing.T) {
  t.Run("testCase", func(t *testing.T) {
  })
}

t.Log("GiventheneedtotesttheSendJSONendpoint.") {
  //code
}

t.Run("Failed to send email", func(t *testing.T) {})
t.Errorf("Got: %v, want: %v", a, e)
t.Skip("Skipping...")
b.RunParallel(func(pb *testing.PB) {})
````

## Bench

````golang
b.ResetTimer()
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
