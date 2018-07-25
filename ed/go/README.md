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

go build                 # Compiles and installs packages and dependencies
go build -race ./example #
go env
go fmt ./...             # Format code
go get ./...             # Install all project dependencies
go golint ./...          # Check code
go install               # Install packages and dependencies
go list                  # List packages
go list ...
go run --work ed/go/examples/hw.go
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

// func as interface
type Func func(fl FieldLevel) bool
type LogFunc func(ctx context.Context, format string, args ...interface{})
// ✳️
type F func (s string) string
func (f F) Bar(s string) {
}

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

v := varI.(T)
if v, ok := varI.(T); ok {
    Process(v) return
}
// varI is not of type T

switch t := areaIntf.(type) {
    case *Square:
        fmt.Printf(“Type Square %T with value %v\n”, t, t)
    case *Circle:
        fmt.Printf(“Type Circle %T with value %v\n”, t, t)
    case float32:
        fmt.Printf(“Type float32 with value %v\n”, t)
    case nil:
        fmt.Println(“nil value: nothing to check?”) default: fmt.Printf(“Unexpected type %T”, t)
}

// non-struct type
type MyFloat float64

// dbg
f, _ := os.OpenFile("/tmp/debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777); f.WriteString("dbg" + "\n")
// or
logfile, _ := os.Create("/tmp/debug.log")
defer logfile.Close()
logger := log.New(logfile, "[example] ", log.LstdFlags|log.Lshortfile)
logger.Println("This is a regular message.")

// Verify statically that *Transport implements http.RoundTripper.
var _ http.RoundTripper = (*Transport)(nil)

// to safe update variable from goroutines
Import "sync/atomic"
atomic.AddInt64(&counter, 1)

log.SetPrefix("TRACE: ")
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
log.Println("message")
````

Go is compiled, garbage-collected, concurrent, type-safe.

A workspace is a directory hierarchy with three directories at its root:
* bin (executable commands)
* pkg (package objects)
* src

Put code into `internal` dir to make it private.

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
* Map (passes by ref)
* Channel (passes by ref)
* Function
* Interface
* Struct
* Pointer (passes by ref)

Ref variables stores in the heap, which is garbage collected
and which is a much larger memory space than the stack.

**array**
An array's size is fixed.
An array variable denotes the entire array; it is not a pointer to the first array element, etc.

An array's length is part of its type, so arrays cannot be resized.

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
(error in case when parameter is not slice but array).

The **append** function appends the elements x to the end of the slice s,
and grows the slice if a greater capacity is needed.

**struct** is a collection of fields.

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

#### Concurrency

Go uses the concurrency model called Communicating Sequential Processes (CSP).

Two crucial concepts make Go’s concurrency model work:
* Goroutines
* Channels

#### Goroutine

Do not communicate by sharing memory. Instead, share memory by communicating.
<br>⚠️ Do not use global variables or shared memory, they make your code unsafe for running concurrently.

`kill -6 {PID}` kill the program and give a stack trace for each goroutine.

The operating system schedules threads to run against processors
regardless of the process they belong to.

The operating system schedules threads to run against physical processors
and the Go runtime schedules goroutines to run against logical processors.

If you want to run goroutines in parallel, you must use more than one logical processor.
But to have true parallelism, you still need to run your program
on a machine with multiple physical processors. 

`runtime.GOMAXPROCS(1)` - tell the scheduler to use a single logical processor for program.

Goroutine is operating on a separate function stack hence no recover from panic,
([img](https://monosnap.com/file/FyeRMIaPfHmuQStwoqBkt4PxWRwSfJ)).

#### Channel

````
c <- 42    // write to a channel
val := <-c // read from a channel

c1 := make(<-chan bool)   // can only read from
c2 := make(chan<- []bool) // can only write to
````
````
ch := make(chan type, value)
# where:
# value == 0 (blocking) synchronous, unbuffered 
# value > 0 (non-blocking) asynchronous, buffered, up to value elements 
````

Blocking channels:

1. A send operation on a channel blocks until a receiver is available.
   No recipient for the value on ch - no other value can be put in the channel

2. A receive operation for a channel blocks until a sender is available.
   If there is no value in the channel, the receiver blocks.

All operations on unbuffered channels block the execution
until both sender and receiver are ready to communicate.
That’s why unbuffered channels are also called synchronous.

In case a channel has a buffer - all read operations succeed without blocking
if the buffer is not empty,
and write operations - if the buffer is not full.
These channels are called asynchronous.

````
for i := 1; i <= 9; i++ {
    select {
    case msg := <-c1:
        println(msg)
    case msg := <-c2:
        println(msg)
    case msg := <-c3:
        println(msg)
    }
}
````

Range chan: range will work until the channel is closed explicitly.

Closing a channel has one more useful feature - reading operations on closed channels
do not block and always return default value for a channel type.

Only the sender should close a channel, never the receiver
(otherwise panic).

#### Memory Management

There are 3 places memory can be allocated:

* the stack - functions parameters, local variables allocated on the stack.
  Each goroutine has its own stack.
  Goroutine stacks are allocated on the heap.
  If the stack needs to grow then heap operations (allocate new, copy old to new, free old) will occur.

* the heap - does not have a single partition of allocated and free regions, set of of free regions.
  Unlike the stack, the heap is not owned by one goroutine
  (manipulating the set of free regions in the heap requires synchronization).

* the data segment - this is where global variables are stored.
  Defined at compile time and therefore does not grow and shrink at runtime.

Escape analysis (variable & pointer analysis when function exit)
is used to determine whether an item can be allocated on the stack.

The Go garbage collector occasionally has to stop the world to complete the collection task.
The stop the world task will take no more than 10 milliseconds
out of every 50 milliseconds of execution time.
