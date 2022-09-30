APP_NAME := "quizengine"
COMMIT_HASH := $(shell git rev-parse HEAD)
SPANNER_EMULATOR_HOST := "localhost:9010"

# Run golang lint.
.PHONY: lint
lint:
	go vet ./...
	golangci-lint run

# Run golang tests.
.PHONY: test
test:
	go test -race -v ./app/...
	go test -race -v ./cmd/...
	go test -race -v ./config/...

# Build golang binary.
.PHONY: build
build: BUILD_VARS:=-ldflags="-X main.BuildCommitHash=$(COMMIT_HASH)"
build: GOOS ?= linux
build: GOARCH ?= amd64
build:
	@echo "==> Starting build"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(APP_NAME) $(BUILD_VARS) ./cmd/$(APP_NAME)
	@chmod +x $(APP_NAME)
	@echo "==> Done! Binary size: "
	@du -h $(APP_NAME)

# Run golang binary.
.PHONY: run
run:
	SPANNER_EMULATOR_HOST=$(SPANNER_EMULATOR_HOST) ./$(APP_NAME) -config-path=config/local.json

# Start Spanner.
.PHONY: spanner-start
spanner-start: CFG ?= to
spanner-start: DBI ?= test-instance
spanner-start: DB ?= test-db
spanner-start:
	@ops/script/spanner.sh spanner-start $(CFG) $(DBI) $(DB)
	@ops/script/spanner.sh spanner-init $(DBI) $(DB)

# Stop Spanner.
.PHONY: spanner-stop
spanner-stop:
	@ops/script/spanner.sh spanner-stop
