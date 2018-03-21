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

* Don't pass pointers as function arguments (just to save a few bytes). 
  This advice does not apply to large structs.

* Variable names in Go should be short rather than long.
