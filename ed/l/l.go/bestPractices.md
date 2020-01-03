Best Practices
-

* Avoid nesting by handling errors first.
* Make your packages "go get"-able.
* Use goroutines to manage state.
* Use separated `doc.go` for pkg documentation.
* Follow:
    * on save: `go fmt ./... && golint ./...` or `imports`
    * on build: `go vet` and `golint` and `go test`
    * on deploy: `go test -tags=integration`
* `camelCase` for constants.
* `snake_case` for filenames.
* `lowercase` for packages, like:
  `"index/suffixarray", "mime/quotedprintable", "net/http/httptrace"`.
* Variable names in Go should be short rather than long.
* Don't pass pointers as function arguments (just to save a few bytes).
  This advice does not apply to large structs.
* When defining methods: use a pointer to a type (struct) as a receiver.
* Function must return error as last value.
* Use `crypto/rand` to generate keys.
* When you spawn goroutines, make it clear when - or whether - they exit.
* Name function's return variables.
* Use `-race` option.

#### Anti-Patterns

* [Tiny package syndrome](https://www.youtube.com/watch?v=ltqV6pDKZD8&feature=youtu.be&t=7m30s).
* Premature Exportation - over export small struct, func, etc.
* Config structs (it's fat like GOD obj).
* Pointer all - bad for GC.
* Async APIs - how to shut down goroutine?
* Panic in a Lib.
* Interface all things.
* Naked `return`.
* `interface{}` (no clue what the data inside).

#### Tricky

* It's better no name return values.
* `pkg.New()` - ok in case you have 1 struct per package,
otherwise have to have: `pkg.NewUser() & pkg.NewLocation()`.
* Inline error check (`if err := x.f(); err != nil {}`) - debugger won't stop before error check.