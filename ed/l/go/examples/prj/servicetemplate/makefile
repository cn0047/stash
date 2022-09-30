APP_NAME := $(shell basename "$(CURDIR)")
COMMIT_HASH := $(shell git rev-parse HEAD)

# Build golang binary.
.PHONY: build
build: BUILD_VARS:=-ldflags="-X main.BuildCommitHash=$(COMMIT_HASH)"
build: GOOS ?= linux
build: GOARCH ?= amd64
build:
	@echo "==> Starting build"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME) $(BUILD_VARS) ./cmd/$(APP_NAME)
	@echo "==> Done! Binary size: "
	@du -h $(APP_NAME)

# Get golang swagger.
.PHONY: get-swagger
get-swagger:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger

# Generate golang code from api.yaml swagger file.
.PHONY: go-swagger
go-swagger: SWAGGER_YAML ?= api.yaml
go-swagger:
	@echo "Swagger version:"
	@${GOPATH}/bin/swagger version
	@rm -rf ./swagger/*
	@${GOPATH}/bin/swagger generate server --exclude-main -s rest -m restmodel \
		-A $(APP_NAME) -f $(SWAGGER_YAML) -t ./swagger
