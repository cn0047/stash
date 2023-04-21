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

# Run golang unit tests.
.PHONY: test-unit
test-unit:
	go test -v ./app/...
	go test -v ./cmd/...
	go test -v ./config/...

# Init and run golang BDD tests.
.PHONY: test-spec
test-spec: spanner-stop spanner-start spanner-load-fixtures test-spec-run

# Run golang BDD tests.
.PHONY: test-spec-run
test-spec-run:
	SPANNER_EMULATOR_HOST=localhost:9010 go test -v ./test/...

# Get golang swagger.
.PHONY: get-swagger
get-swagger:
	@go get -u github.com/go-swagger/go-swagger/cmd/swagger

# Generate golang code from api.yaml swagger file.
.PHONY: go-swagger
go-swagger: SWAGGER_YAML ?= api.yaml
go-swagger: SWAGGER_JSON ?= app/swagger/apidocs/api.json
go-swagger:
	@echo "Swagger version:"
	@${GOPATH}/bin/swagger version
	@rm -rf ./app/swagger/rest
	@rm -rf ./app/swagger/restmodel
	@${GOPATH}/bin/swagger generate server --exclude-main -s rest -m restmodel \
		-A $(APP_NAME) -f $(SWAGGER_YAML) -t ./app/swagger
	@${GOPATH}/bin/swagger generate spec -i $(SWAGGER_YAML) -o $(SWAGGER_JSON)

# Start Spanner.
.PHONY: spanner-start
spanner-start: CFG ?= to
spanner-start: DBI ?= test-instance
spanner-start: DB ?= test-db
spanner-start:
	@.ops/script/spanner.sh spanner-start $(CFG) $(DBI) $(DB)

# Load test fixtures for Spanner.
.PHONY: spanner-load-fixtures
spanner-load-fixtures: DBI ?= test-instance
spanner-load-fixtures: DB ?= test-db
spanner-load-fixtures:
	@.ops/script/spanner.sh spanner-load-fixtures $(DBI) $(DB)

# Stop Spanner.
.PHONY: spanner-stop
spanner-stop:
	@.ops/script/spanner.sh spanner-stop
