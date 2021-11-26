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
