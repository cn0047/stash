GO
-

go1.9.3

[online editor](https://play.golang.org/)
[packages](https://golang.org/pkg/) [and](https://godoc.org/)
[examples](https://gobyexample.com/)
[badges](https://goreportcard.com/)

````
$GOROOT
$GOPATH     // environment variable specifies the location of your workspace.
$GOMAXPROCS // number of OS threads that can execute user-level Go code simultaneously.

go run --work ed/go/examples/hw.go

# Install all project dependencies
go get ./...

# Format code
go fmt ./...

# Check code
go golint ./...

go vet
````

````go
import (
    "fmt"
    "os"
)

var (
    msq = "str"
    message string = "str"
)
x := 1
var id = 1
var power int
power = 9000
var power2 int = 9000
fmt.Printf("It's over %d\n", power)

const (
    PI = 3.14
    A = iota
    B = iota
    C = iota
)
const s string = "constant"
const n = 500000000

fmt.Printf("Type: %T\n", myType)
fmt.Printf("%+v", myType)
fmt.Printf("%#v", myType)

// init module
func init() {}

if p := "MR"; isFormal {
}

if len(os.Args) != 2 {
    os.Exit(1)
}
fmt.Println("It's over", os.Args[1])

// read input from CLI
var option string
fmt.ScanIn(&option)
println(option)

for i := 0; i < 100; i++ {
}
````

Go is compiled, garbage-collected, concurrent, type-safe.

A workspace is a directory hierarchy with three directories at its root:
* bin (executable commands)
* pkg (package objects)
* src

**new** keyword - allocates memory, does not initialize the memory
and returns its address (pointer).

**make** - creates slices, maps, and channels
and it returns an initialized (not zeroed) value of type T (not `*T`).
The reason for the distinction with new keyword - is that these three types represent,
under the covers, references to data structures that must be initialized before use.
<br>A slice, for example, is a three-item descriptor
containing a pointer to the data (inside an array), the length, and the capacity,
and until those items are initialized, the slice is nil.

**defer** pushes a function call in list,
list will be executed after the surrounding function returns.
Example: `defer file.Close()`.

**panic** is a built-in function that stops the ordinary flow of control and begins *panicking*.

**recover** is a built-in function that regains control of a panicking goroutine.

**copy** - copies elements from a source slice into a destination slice.

OOP in go represented by **structures**. Inheritance in go represented by composition.

#### Data types

Basic types:

* bool
* string
* int (int8, int16, int32, int64)
* uint (uint8, uint16, uint32, uint64, uintptr)
* byte // alias for uint8
* rune // alias for int32
* float32 float64
* complex64 complex128

Other types:

* Array
* Slice
* Struct
* Pointer
* Function
* Interface
* Map (passes by ref)
* Channel

#### Interface

````go
type Logger interface {
    Log(message string)
}
````

#### HTTP

````
url := req.URL
query := url.Query()
q := query["q"]
page := query.Get("page")

err := req.ParseForm()
f := req.Form
un := f["username"]
p := f.Get["username"]
````

#### Goroutine

#### Channel
