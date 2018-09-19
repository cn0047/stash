modules
-

[doc](https://github.com/golang/go/wiki/Modules)

Module directory should be outside your $GOPATH
because by default, the modules support is disabled inside it.

````
export GO111MODULE=on

go mod init
go mod init example.com/my/module/v2

go build ./...
go test all
````
