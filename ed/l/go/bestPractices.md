Best Practices
-

[code style](https://github.com/golang/go/wiki/CodeReviewComments)
[code style](https://google.github.io/styleguide/go)
[code style](https://google.github.io/styleguide/go/best-practices)
[uber code style](https://github.com/uber-go/guide/blob/master/style.md)
[project layout](https://github.com/golang-standards/project-layout)
[recipes](https://github.com/nikolaydubina/go-recipes)
[panic detector](https://github.com/uber-go/nilaway/)
[package import order](https://github.com/daixiang0/gci)
[linter](https://github.com/golangci/golangci-lint)
[linter](https://github.com/dominikh/go-tools)
[error linter](https://github.com/kisielk/errcheck)
[raft protocol](https://github.com/hashicorp/raft)

* Avoid nesting by handling errors first.
* Make your packages "go get"-able.
* Use goroutines to manage state.
* Use separated `doc.go` for pkg documentation.
* Use `-race` option.
* Follow:
    * on save: `go fmt ./... && golint ./...` or `imports`.
    * on build: `go vet && golint && go test`.
    * on deploy: `go test -tags=integration`.
* `camelCase` for constants.
* `snake_case` for filenames.
* `lowercase` for packages (directories), like: `"net/http/httptrace"`.
* Variable names in Go should be short rather than long.
* When defining methods: use a pointer to a type (struct) as a receiver.
* Function must return error as last value.
* Name function's return variables.
* When you spawn goroutines, make it clear when they exit.

* Use `crypto/rand` to generate keys.
* Don't pass pointers as function arguments (just to save a few bytes).
  This advice does not apply to large structs.
* Always close the response body `defer resp.Body.Close()`.
* Always read response body even the data is not important
  `_, err = io.Copy(ioutil.Discard, resp.Body)`
  to avoid memory leak when reusing http connection.
* Don't forget about padding bytes in structures declaration (8 byte).

#### Anti-patterns

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

* It's better to name return values.
* `pkg.New()` - ok in case you have 1 struct per package,
otherwise have to have: `pkg.NewUser(); pkg.NewLocation()`.
* Inline error check (`if err := x.f(); err != nil {}`) - debugger won't stop before error check.
* Use underscore in test cases names (`ts.Run("test_case", ...`).

#### Non-classified

* Ternary operator `$cnd ? true : false;`.
* `iota` (have to calculate value every time, or rely on IDE).
* No `this, self`.
* `*_test.go` for files, but nothing for directories (mocks, fixtures, testdata).
