Test
-

## Test

````sh
# to get more test info
go test -gcflags
go build -gcflags '-m'
go build -gcflags '-m -m' ./main.go

go test -json

goapp test -v transactional/users/service -run TestGetServiceAccountsForAdmin
````

````go
t.Log("GiventheneedtotesttheSendJSONendpoint.") {
  //code
}

t.Run("Failed to send email", func(t *testing.T) {})
t.Errorf("Got: %v, want: %v", a, e)
t.Skip("Skipping...")
b.RunParallel(func(pb *testing.PB) {})
````

## Bench

````go
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
