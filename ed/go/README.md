GO
-

go1.9.3

[online editor](https://play.golang.org/)
[packages](https://golang.org/pkg/) [and](https://godoc.org/) [and](https://gopkg.in/)
[examples](https://gobyexample.com/)
[badges](https://goreportcard.com/) [covererage](https://gocover.io/)

````
$GOROOT       // root of the go (`/usr/local/go/`)
$GOROOT_FINAL // if go moved from $GOROOT
$GOPATH       // environment variable specifies the location of your workspace.
$GOMAXPROCS   // number of OS threads that can execute user-level Go code simultaneously.

export GOROOT=$HOME/go
export GOBIN=$GOROOT/bin
export GOARCH=amd64
export GOOS=linux

go run --work ed/go/examples/hw.go

# Compiles and installs packages and dependencies
go build

# Install packages and dependencies
go install

# Install all project dependencies
go get ./...

# Format code
go fmt ./...

# Check code
go golint ./...

go vet
````

````
# Fix finds Go programs that use old APIs and rewrites them to use newer ones
go tool fix

# Performs precise type-safe renaming of identifiers
gorename

# moves go packages, updating import declarations
gomvpkg
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
    Monday, Tuesday, Wednesday = 1, 2, 3
)
const s string = "constant"
const n = 500000000

# https://golang.org/pkg/fmt/#hdr-Printing
fmt.Printf("Type: %T\n", myType)
fmt.Printf("%+v", myType)
fmt.Printf("%#v", myType)

// init module
func init() {}

type Func func(fl FieldLevel) bool

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

**make** - creates only slices, maps, and channels
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

**OOP** in go represented by **structures**. Inheritance in go represented by composition.

#### Data types

Basic types:

* bool
* string
* int (int8 (aka byte), int16, int32 (aka rune), int64)
* uint (uint8, uint16, uint32, uint64, uintptr (large enough to hold the bit pattern of any pointer))
* float32 float64
* complex64 complex128

Other types:

* Array
* Slice (passes by ref)
* Struct
* Pointer (passes by ref)
* Function
* Interface
* Map (passes by ref)
* Channel (passes by ref)

Ref variables stores in the heap, which is garbage collected
and which is a much larger memory space than the stack.

**array**
An array's size is fixed.
An array variable denotes the entire array; it is not a pointer to the first array element, etc.

**slice**
Slice is a descriptor of an array segment.
It consists of a pointer to the array, the length of the segment, and its capacity.
The length is the number of elements referred to by the slice.
The capacity is the number of elements in the underlying array.
<br>The zero value of a slice is nil.
The len and cap functions will both return 0 for a nil slice.

`make` allocates an array and returns a slice that refers to that array.

Slicing does not copy the slice's data.
It creates a new slice value that points to the original array.
Therefore, modifying the elements (not the slice itself) of a re-slice
modifies the elements of the original slice.

To increase the capacity of a slice one must create a new, larger slice
and copy the contents of the original slice into it.

**copy** - copies elements from a source slice into a destination slice.

The **append** function appends the elements x to the end of the slice s,
and grows the slice if a greater capacity is needed.

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

Do not communicate by sharing memory. Instead, share memory by communicating.
<br>⚠️ Do not use global variables or shared memory, they make your code unsafe for running concurrently.

#### Channel
