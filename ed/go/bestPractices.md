Best Practices
-

* Avoid nesting by handling errors first.
* Make your packages "go get"-able.
* Use goroutines to manage state.
* Use separated `doc.go` for pkg documentation.
* Follow:
    * on save: `go fmt` or `imports`
    * on build: `go vet` and `golint` and `go test`
    * on deploy: `go test -tags=integration`
