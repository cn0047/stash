Test
-

Bench output:

````
goos: linux
goarch: amd64
BenchmarkFib1-2             100000000           11.0 ns/op         0 B/op          0 allocs/op
BenchmarkFib5-2             20000000            95.0 ns/op         0 B/op          0 allocs/op
BenchmarkFib10and20/10-2     1000000          1069 ns/op           0 B/op          0 allocs/op
BenchmarkFib10and20/20-2       10000        132025 ns/op           0 B/op          0 allocs/op
BenchmarkFibComplete-2       1000000          1075 ns/op           0 B/op          0 allocs/op
PASS
ok      _/gh/ed/go/examples/bench   7.594s
````

Means that the loop ran 100000000 times at a speed of 11.0 ns (nanosecond) per loop.
