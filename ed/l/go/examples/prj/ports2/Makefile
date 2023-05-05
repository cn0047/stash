APP_NAME := $(shell basename "$(CURDIR)")

.PHONY: build
build: GOOS ?= linux
build: GOARCH ?= amd64
build:
	@echo "==> Starting build"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME) ./cmd/$(APP_NAME)
	@echo "==> Done! Binary size: "
	@du -h $(APP_NAME)

.PHONY: build-docker
build-docker:
	@echo "==> Starting build docker image"
	docker build -t $(APP_NAME):latest .
	@echo "==> Done!"

.PHONY: run
run: build build-docker
	@echo "==> Starting application"
	@docker-compose up -d
	@echo "==> Application started."

.PHONY: check
check:
	@go vet ./...
	@go fmt ./...
	@golangci-lint run -c .golangci.yaml

.PHONY: test
test:
	@go test -v ./app/...
	@go test -v ./config/...
