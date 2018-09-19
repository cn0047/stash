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

* Don't pass pointers as function arguments (just to save a few bytes).
  This advice does not apply to large structs.

* Variable names in Go should be short rather than long.

* When defining methods: use a pointer to a type (struct) as a receiver.

* Function must return error as last value.

* Use `crypto/rand` to generate keys.

* When you spawn goroutines, make it clear when - or whether - they exit.
