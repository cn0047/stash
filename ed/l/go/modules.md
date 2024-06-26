modules
-

[doc](https://github.com/golang/go/wiki/Modules)
[dependency analysis](https://github.com/loov/goda)

A module is a collection of related Go packages
that are versioned together as a single unit.

Modules must be semantically versioned according to semver,
such as `v0.1.0, v1.2.3, or v1.5.0-rc.1`. The leading v is required.

Module directory should be outside your $GOPATH
because by default, the modules support is disabled inside it.

Modules won't work with relative imports like `import "./subdir"`.

````sh
# for go.mod inside GOPATH
export GO111MODULE=on

# specifies private modules
export GOPRIVATE=*.corp.example.com,rsc.io/private
export GOPRIVATE=github.com/prvtOrg
export GOFLAGS=-mod=readonly
export GOFLAGS=-mod=vendor
export GOFLAGS=-mod=

# to fix error: fatal: could not read Username for 'https://*.com': terminal prompts disabled
git config --global url."git@github.com:".insteadOf "https://github.com/"
git config --global url."git@bitbucket.org:".insteadOf "https://bitbucket.org/"
git config --global url."git@github.com:org/".insteadOf "https://github.com/org/"

go mod init
go mod init example.com/my/module/v2

go build ./...
go test all

go get google.golang.org/appengine@'>=v1.2.0'

go list -m
go list -m -json all
go list -m all # view final versions that will be used in a build
go list -u -m all # view available minor and patch upgrades
go list -m -versions rsc.io/sampler
go list -deps -f '{{with .Module}}{{.Path}} {{.Version}}{{end}}' ./... | sort -u
go list -f '{{range $imp := .Imports}}{{printf "%s\n" $imp}}{{end}}' | sort # imports for current directory package

go get foo@v1.2.3 # get specific versions of pkg

go mod tidy # prune any no-longer-needed dependencies (and add any dependencies needed)
go mod why -m <module>
go mod graph

go clean -modcache # ✅
go clean --modcache

# download modules to local cache
go mod download
# create a vendor dir and copy all dependencies into it
go mod vendor

go doc rsc.io/quote/v3

go get -d github.com/thepkg/strings # get not to build or install package, just download
# update all ✅
go get -u ./... or go get -u=patch ./... # update all direct and indirect dependencies
go get -u all
go get -u github.com/thepkg/strings@master
go mod tidy
go mod download
go mod vendor
````

Typical day-to-day workflow can be:
add import statements to your .go code as needed,
standard commands like go build or go test will automatically add new dependencies.

`go.mod`:

````sh
module github.com/my/thing
require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency/v4 v4.0.0
)
````

`go.sum`:

Contains the expected cryptographic checksums of the content of specific module versions.
This file must be commited into git.

`vendor/modules.txt` describes how each entry in go.mod file corresponds to which subdirectory within `vendor/`.

`go.work` for multiple modules workspace (file syntax is similar to go.mod file).

````sh
go work init

go work use $m # add module to workspace

go work use -r # recursively add/remove directories to workspace

go work sync
````
