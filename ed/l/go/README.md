Go (Golang)
-
<br>1.13
<br>1.10.1
<br>1.9.3
Since 2009.

[doc](https://golang.org/)
[github](https://github.com/golang/go/wiki)
[spec](https://golang.org/ref/spec)
[cmd](https://golang.org/cmd/go/)
[packages](https://golang.org/pkg/) [and](https://godoc.org/) [and](https://gopkg.in/)
[golang.org/x/\*](https://pkg.go.dev/search?q=golang.org%2Fx)
[examples](https://gobyexample.com/)
[online editor](https://play.golang.org/) [and](https://goplay.space/)
[sizeof](http://golang-sizeof.tips/)
[badges](https://goreportcard.com/)
[avatar](https://gopherize.me/)
[pprof-ui](https://www.speedscope.app/)
[city](https://go-city.github.io/#/github.com/cn007b/monitoring)
[DIC (wire)](https://github.com/google/go-cloud/tree/master/wire)
[DIC (wire) tutorial](https://github.com/google/go-cloud/tree/master/samples/wire)
[load test](https://github.com/tsenart/vegeta)
[fultext search](https://github.com/blevesearch/bleve)
[distributed tracing (zipkin-go)](https://zipkin.io/)
[protocol for microservices](https://github.com/uber/tchannel)
[golangci](https://golangci.com/)
[security](https://github.com/securego/gosec)
[covererage](https://gocover.io/) [![and](https://gocover.io/_badge/github.com/thepkg/strings)](https://gocover.io/github.com/thepkg/strings)
[json to struct](https://mholt.github.io/json-to-go/)
[sql code gen](https://github.com/kyleconroy/sqlc)

Go - statically-typed, general-purpose, with type safety
and automatic memory management programming language.

````bash
export PATH=$PATH:/Users/kovpakvolodymyr/go/bin

$GOROOT       # root of the go (`/usr/local/go/`)
$GOROOT_FINAL # if go moved from $GOROOT
$GOPATH       # environment variable specifies the location of your workspace.
$GOBIN        # GOBIN=$GOROOT/bin
$GOCACHE      #
$GOOS         # (windows|linux)
$GOARCH       # (386|amd64)
$GOMAXPROCS   # number of OS threads that can execute user-level Go code simultaneously.
$GOTRACEBACK  # (none|single|all|system|crash) verbosity level in case of panic
$GOGC         # Garbage Collector ↓
$GODEBUG      # https://godoc.org/runtime#hdr-Environment_Variables
              # scheddetail=1 - detailed scheduler info
              # schedtrace=1000 - single line to standard error every 1000 milliseconds
              # gctrace=1 - single line to standard error at each collection
              # GODEBUG=allocfreetrace=1,gcpacertrace=1,gctrace=1,scavenge=1,scheddetail=1,scheddetail=1,schedtrace=1000
              # GODEBUG=gocacheverify=1,gocachehash=1,gocachetest=1

export GOROOT=$HOME/go
export GOFLAGS=-mod=vendor

GOARCH=386 go build
go build                     # compiles and installs packages and dependencies
go build -race ./example     #
go tool compile -S x.go      # assembly output
go build -gcflags='-v'       # verbose
go build -gcflags '-S' x.go  # assembly output
go build -gcflags='-m'       # print optimization decisions
go build -gcflags='-m -m'    # -m - escape analysis
go build -gcflags='-m -l'    # -l - omit inlining decisions
go build -gcflags='-m -m -l' #
go build -mod=vendor
go build -o newname
GOOS=linux GOARCH=amd64 go build \
  -ldflags="-w -s" -o /go/bin/hello # build in docker (get the smallest binaries)
                                    # -w - not include info for debug, pprof, etc (DWARF).
                                    # -s - disable symbol table, `go tool nm`
  -ldflags="-X github.com/org/repo/pkgName.BuildCommit=$GO_BUILD_COMMIT"

go clean -r                           # clean unneeded files
go env
go env GOOS
go env -json
go env GOPATH
go fmt ./...                          # format code
go get ./...                          # install all project dependencies
go golint ./...                       # check code
go install                            # install packages and dependencies
go list                               # list packages
go list -f '{{ join .Imports "\n" }}' # directly imported packages
go list -f '{{ join .Deps "\n" }}'    # all subdependencies
go list -json -f {{.Deps}} net/http
go list -e -json -f {{.Deps}} products/service/scraper
go list ...
go run --work ed/go/examples/whatever/hw.go # see the location of temporary exec file
go run -race ed/go/examples/whatever/hw.go
go vet # examines Go coge, reports suspicious constructs (Printf with wrong arguments).
go vet -all ./...
go bug # creates issue on github
godoc -http=:6060 -goroot=$PWD
go tool fix # finds Go programs that use old APIs and rewrites them to use newer ones
gorename # Performs precise type-safe renaming of identifiers
gomvpkg # moves go packages, updating import declarations

# Gopkg.toml, Gopkg.lock
dep init
dep status
dep ensure         # ensure a dependency is safely vendored in the project
dep ensure -update # update the locked versions of all dependencies

gosec -tests ed/go/examples/bench/...
gosec -fmt=json -out=/tmp/gosec_output.json -tests ed/go/examples/bench/...

strace go_bin_file
ptrace
````

````golang
// +build tools
// ↑ ⚠️ build tag comment: ensures file.go won't be compiled into the binary

import (
    "fmt"
    "os"
    // when file git.x.com/go/prj/pkg/file/file_executor.go
    // has package filexecutor
    file_executor "git.x.com/go/prj/pkg/file"
)

var (
    msg = "str"
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

sort.Slice(arrInt32, func(i, j int) bool { return arrInt32[i] < arrInt32[j] })
sort.Strings(strSlice)

// https://golang.org/pkg/fmt/#hdr-Printing
fmt.Printf("%T", myType)
fmt.Printf("%t", myBool)
fmt.Printf("%+v", myVal)
fmt.Printf("%p", myPointer)

// init pkg
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
for c := n.FirstChild; c != nil; c = c.NextSibling {
  f(c)
}

v := varI.(T)
if v, ok := varI.(T); ok { // "comma ok" idiom.
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
        fmt.Println(“nil value: nothing to check?”)
    default: fmt.Printf(“Unexpected type %T”, t)
}

// non-struct type
type MyFloat float64

// debug
f, _ := os.OpenFile("/tmp/debug.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777); f.WriteString(fmt.Sprintf("%+v \n", 1))
// or
logfile, _ := os.Create("/tmp/debug.log")
defer logfile.Close()
logger := log.New(logfile, "[example] ", log.LstdFlags|log.Lshortfile)
logger.Println("This is a regular message.")

// Verify statically that *Transport implements http.RoundTripper.
var _ http.RoundTripper = (*Transport)(nil)

// check that obj has method
if correctobj, ok := obj.(interface{methodName()}); ok {
  correctobj.methodName()
}

log.SetPrefix("TRACE: ")
log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
log.Println("message")

"html/template"
template.HTMLEscapeString(input)
template.JSEscapeString(input)
````

#### HTTP

````golang
http.Server{IdleTimeout}

b, err := httputil.DumpRequest(ctx.Request(), true)

query := req.URL.Query()
q := query["q"]
page := query.Get("page")
req.FormValue("timeout")

err := req.ParseForm()
f := req.Form
un := f["username"] // array of values
p := f.Get["username"] // value

tr := &http.Transport{DisableKeepAlives: true}
client := http.Client{Transport: tr}
client := http.Client{Timeout: time.Millisecond * timeout}

// by default behaves like curl -L
http.Client{}.Get(URL)

// w http.ResponseWriter
w.Write([]byte("html"))

// to enable http2
err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)

// Handle vs HandleFunc
http.Handle("/img/", http.FileServer(http.Dir("public")))
mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
````

Go is compiled, garbage-collected, concurrent, type-safe.
Go source code is always UTF-8.

A workspace is a directory hierarchy with three directories at its root:
* bin (executable commands)
* pkg (package objects)
* src (source code)

Put code into `internal` dir to make it private (magic).

Sentinel error -  custom error value
(standard library: sql.ErrNoRows, io.EOF, etc).
[wrap error tool](github.com/pkg/errors)
<br>`errors.Is` like a comparison to a sentinel error,
<br>`errors.As` like type assertion (`if e, ok := err.(*QueryError); ok { … }`).
<br>`fmt.Errorf("decompress %v: %w", name, err)`

**Context** carries deadlines, cancelation signals, and other request-scoped values
across API boundaries and between processes.
Ctx may be: WithCancel, WithDeadline, WithTimeout, or WithValue.

#### Data types

Basic types:
* bool
* string
* int (int8 (aka byte), int16, int32 (aka rune), int64)
* uint (uint8, uint16, uint32, uint64, uintptr (large enough to hold the bit pattern of any pointer))
* float32, float64
* complex64, complex128

Other types:
* array
* slice (passes by ref)
* map (passes by ref)
* channel (passes by ref)
* function
* interface
* struct
* pointer (*passes by ref*)

Comparable types:
* bool
* string
* int
* float
* complex

* array (if values of the array element type are comparable)
- slice (NOT COMPARABLE)
- map (NOT COMPARABLE)
* channel (if they created by the same call to `make` or nil)
- func (NOT COMPARABLE)
* interface (if they have identical dynamic types and equal dynamic values or nil)
* struct (if all their fields are comparable)
* pointer (if they point to the same var or nil)

Ref variables stores in the heap, which is garbage collected
and which is a much larger memory space than the stack.

**new** keyword - allocates memory, does not initialize the memory
and returns its address (pointer).

**make** - creates only slices, maps, and channels
and it returns an initialized (not zeroed) value of type T (not `*T`).
The reason for the distinction with new keyword - is that these three types represent,
under the covers, references to data structures that must be initialized before use.
<br>
`make([]int, 0, 10) // slice of length 0 and capacity 10.`

**defer** pushes a function call in list,
list will be executed after the surrounding function returns.
Example: `defer file.Close()`.

**panic** is a built-in function that stops the ordinary flow of control and begins *panicking*.

**recover** is a built-in function that regains control of a panicking goroutine.

**range** (`for name, value := range m {}`) over map - there is no guarantee
regarding the order of hash map.

**string**
String it is immutable slice of bytes.
When we store a character value in a string, we store its byte-at-a-time representation.
Go strings aren’t always UTF-8, only string literals are UTF-8.
String values can contain arbitrary bytes
(string literals always contain UTF-8 text as long as they have no byte-level escapes).

**map**

`Key` may be any type that is comparable and `Value` may be any type at all.

`make(map[int]int, 100)`
the initial capacity does not bound its size, it's only a hint.
The goal to use this capacity - is to not resize map to often.
If you know that you will store 100 entries in map,
it's better to allocate map with such (or greater) capacity.

**delete**
Func to delete element from map.

**array**
An array's size is fixed.
An array variable denotes the entire array; it is not a pointer to the first array element, etc.

An array's length is part of its type, so arrays cannot be resized.

**slice**
Slice is a descriptor of an array segment.
It consists of a pointer to the array (pointer to the data inside an array),
the length of the segment, and its capacity
(three-item descriptor, and until those items are initialized, the slice is nil).
The length is the number of elements referred to by the slice.
The capacity is the number of elements in the underlying array.
<br>The zero value of a slice is nil.
The len and cap functions will both return 0 for a nil slice.
<br>
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
<br>
Grows slice by doubling it’s capacity only up to 1024.
After that it will use so-called `memory size classes` to guarantee
that growth will be no more than ~12.5%
<br>
Use append only to append new value to given slice, not to create new slice
(never call append without assigning back to the same variable)!

**OOP** in go represented by **structures**. Inheritance in go represented by composition.

**struct** is a collection of fields.
````
OOP ⇒ class  ⇒ instance ⇒ property
Go  ⇒ struct ⇒ object   ⇒ field
````

**pointer**
[Pointers make stack allocation mostly infeasible](https://segment.com/blog/allocation-efficiency-in-high-performance-go-services/).

#### Interface

Interface — set of methods required to implement such interface.

Interface type (interface{}) — variable of interface type which can hold any value,
meaning no guarantee of any methods at all: it could contain any type.
The comma-ok assignment asks whether it is possible
to convert interface value to some type.
<br>
Interface values are represented as a two-word pair:
the concrete value assigned to the interface variable, and that value's type descriptor.

````golang
type Logger interface {
    Log(message string)
}
````

#### Concurrency

Go uses the concurrency model called Communicating Sequential Processes (CSP).
Two crucial concepts make Go’s concurrency model work: Goroutines & Channels.

#### Goroutine

States:
* not runnable
* runnable
* running

Goroutine - lightweight version of thread, with very low cost of starting up.
Each goroutine is described by struct called G.
Runtime keeps track of each G and maps them onto Logical Processors (P).
P - abstract resource, which needs to be acquired, so OS thread (called M, or Machine) can execute G.
Each P maintains a queue of runnable G‘s.
<br>
When schedule a new goroutine - it's placed into P‘s queue.
<br>
Blocking syscall (opening a file, etc. or network call or channel operations or primitives in the sync package)
G will be intercepted, if there are Gs to run in queue,
runtime will detach the thread from the P (and from M) and create a new thread
(if idle thread doesn’t exist) to service that processor.
When a system calls resumes, the goroutine is placed back.

‼️ Do not communicate by sharing memory. Instead, share memory by communicating.
<br>⚠️ Do not use global variables or shared memory, they make your code unsafe for running concurrently.

`kill -6 {PID}` kill the program and give a stack trace for each goroutine.

Goroutine is operating on a separate function stack hence no recover from panic,
([img](https://monosnap.com/file/FyeRMIaPfHmuQStwoqBkt4PxWRwSfJ)).

Goroutines exists only in the virtual space of go runtime and not in the OS.
Hence, a Go Runtime scheduler is needed which manages their lifecycle.

#### Channel

````golang
c <- 42        // write to a channel
val, ok := <-c // read from a channel

c1 := make(<-chan bool)   // can only read from
c2 := make(chan<- []bool) // can only write to
````
````golang
ch := make(chan type, value)
// where:
// value == 0 (blocking) synchronous, unbuffered
// value > 0 (non-blocking) asynchronous, buffered, up to value elements

// len(c) // count of elements in channel

// make(chan int) equals to make(chan int, 0) equals to make(chan int, 1)
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

````golang
select {
case msg := <-c1:
    println(msg)
case msg := <-c2:
    println(msg)
case msg := <-c3:
    println(msg)
default:
    println(".")
}
````

Range chan: range will work (loop won't stop) until the channel is closed explicitly.
Once chan is closed range loop exit.

Closing a channel has one more useful feature - reading operations on closed channels
do not block and always return default value for a channel type.

Only the sender should close a channel, never the receiver
(otherwise panic).

For standard Go compiler, size of channel element types must be smaller than 65536.

#### Runtime

runtime:
* memory allocation
* channel creation
* goroutines creation (and control goroutines)
* garbage collection

Runtime manages: scheduling, garbage collection, and the runtime environment for goroutines.

Runtime scheduler creates threads (Logical Processors, named P)
over Physical CPUs (called M, or Machine).
`runtime.GOMAXPROCS(numLogicalProcessors)`
`runtime.GOMAXPROCS(1)` - tell the scheduler to use a single Logical Processor for program.
`runtime.GOMAXPROCS` limits number of threads that can execute user-level Go code simultaneously.
For each P (thread) we have goroutines queue.

The OS schedules threads to run against processors regardless of the process they belong to.
FYI: OS thereads expensive to start, stop, utilize.
<br>
The OS schedules threads to run against physical processors
and the Go runtime schedules goroutines to run against logical processors.
<br>
If you want to run goroutines in parallel, you must use more than one logical processor.
But to have true parallelism, you still need to run your program
on a machine with multiple physical processors.

#### GC

During the mark phase, the runtime will traverse all the objects
that the application has references to on the heap and mark them as still in use.
This set of objects is known as live memory.
After this phase, everything else on the heap that is not marked is considered garbage,
and during the sweep phase, will be deallocated by the sweeper.

The Go garbage collector occasionally has to stop the world (pause time) to complete the collection task.
The stop the world task will take no more than 10 milliseconds
out of every 50 milliseconds of execution time.
<br>
The Go GC uses a pacer to determine when to trigger the next GC cycle.
Go’s default pacer will try to trigger a GC cycle every time the heap size doubles.

````sh
$GOGC # Garbage Collector:
      # GOGC=off - disables the gc entirely
      # GOGC=100 - default, means the gc will run each time the live heap doubles.
      # GOGC=200 - will delay start of a gc cycle
      # until the live heap has grown to 200% of the previous size (live set).
````

`GCPercent` (runtime.SetGCPercent) - adjusts how much CPU you want to use and how much memory you want to use.
The default is 100 which means that half the heap is dedicated to live memory and half the heap
is dedicated to allocation.

`MaxHeap` is not yet released but is being used and evaluated internally,
lets the programmer set what the maximum heap size should be.

#### Memory Management

Execution stack for goroutine = 2Kb.

There are 3 places memory can be allocated:
* the stack - function's parameters and local variables allocated on the stack (default maximum stack size ≈1GB limit).
  Each goroutine has its own stack.
  The initial stack size of each goroutine is small (about 2k bytes on 64-bit systems).
  Goroutine stacks are allocated on the heap (‼️).
  If the stack needs to grow then heap operations (allocate new, copy old to new, free old) will occur.
* the heap - does not have a single partition of allocated and free regions, set of of free regions.
  Unlike the stack, the heap is not owned by one goroutine
  (manipulating the set of free regions in the heap requires synchronization).
* the data segment - this is where global variables are stored.
  Defined at compile time and therefore does not grow and shrink at runtime.

Escape analysis (variable & pointer analysis when function exit)
is used to determine whether an item can be allocated on the stack.

Go memory allocator based on TCMalloc memory allocator.
Likewise, TCMalloc Go also divides Memory Pages into block of 67 different classes Size.

**mspan** - (to manage memory pages) it's double linked list object,
with start address of the page, span class, number of pages it contains.

**mcache** (aka Local Thread Cache of Memory) provides memory to goroutine without any locks.
mcache contains a mspan of all class size as a cache.
Object size <=32K byte are allocated directly to mcache using the corresponding size class mspan.
mcache has no free slot - new mspan is obtained from the mcentral.
Size between 16B ~ 32k, calculate the sizeClass to be used and then use the block allocation
of the corresponding sizeClass in mcache.

**mcentral** - is two lists of mspans: empty, nonempty.
mcentral structure is maintained inside the mheap structure.

**mheap** - is the Object that manages the heap in Go, only one global.
mheap has an array of mcentral (this array contains mcentral of each span class).

Object of Size > 32K, is a large object, allocated directly from mheap.
These large request comes at an expenses of central lock, so only one goroutine request
can be served at any given point of time.

#### TCMalloc

The core idea of TCMalloc (thread cache malloc)
is to divide the memory into multiple levels to reduce the granularity of the lock.

Inside TCMalloc memory management is divided into two parts:
  * thread memory:
    each memory page divided into Free List of multiple fixed allocatable size-classes;
    each thread will have a cache for small objects without locks,
    which makes it very efficient to allocate small objects (<=32K) under parallel programs;
  * page heap:
    when allocated Object is larger than 32K, Pages Heap is used for allocation;
    if not enough memory to allocate small objects - go to page heap for memory;
    if not enough - page heap will ask more memory from the Operating System;
