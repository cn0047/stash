modules
-

[doc](https://github.com/golang/go/wiki/Modules)

Module directory should be outside your $GOPATH
because by default, the modules support is disabled inside it.

````sh
# for go.mod inside GOPATH
export GO111MODULE=on

go mod init
go mod init example.com/my/module/v2

go build ./...
go test all

go get google.golang.org/appengine@'>=v1.2.0'

go list -m all
go list -m
````

go.mod:

````
module github.com/my/thing
require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency/v4 v4.0.0
)
````

go.sum:

Contains the expected cryptographic checksums of the content of specific module versions.
