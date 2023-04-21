SPANNER_EMULATOR_HOST=localhost:9010

.PHONY: audit
audit:
	@echo 'Tidy...'
	go mod tidy
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Linting code...'
	golangci-lint run -c .golangci.yaml

# Build golang binary.
.PHONY: build
build: GOOS ?= linux
build: GOARCH ?= amd64
build:
	@echo "==> Starting build"
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o wp ./cmd/wp/main.go
	@echo "==> Done! Binary size: "
	@du -h wp

.PHONY: test
test:
	go test -race ./...

.PHONY: run
run:
	go run cmd/wp/main.go

# Init Spanner.
.PHONY: spanner-init
spanner-init: CFG ?= wp-spanner-emulator
spanner-init: PRJ ?= test-project
spanner-init: DBI ?= outbound-us
spanner-init: DB ?= wp-dev
spanner-init:
	@./main.sh spanner-init $(CFG) $(PRJ) $(DBI) $(DB)

# Init gcloud.
.PHONY: gcloud-init
gcloud-init: CFG ?= wp-spanner-emulator
gcloud-init: PRJ ?= test-project
gcloud-init:
	@./main.sh gcloud-init $(CFG) $(PRJ)

# Load DB migrations into Spanner.
.PHONY: load-migrations
load-migrations: CFG ?= wp-spanner-emulator
load-migrations: PRJ ?= test-project
load-migrations: DBH ?= $(SPANNER_EMULATOR_HOST)
load-migrations: DBI ?= outbound-us
load-migrations: DB ?= wp-dev
load-migrations:
	@./main.sh load-migrations $(CFG) $(PRJ) $(DBH) $(DBI) $(DB)

# Preforms 1 DB migration rollout (up).
.PHONY: up-migration
up-migration: CFG ?= wp-spanner-emulator
up-migration: PRJ ?= test-project
up-migration: DBH ?= $(SPANNER_EMULATOR_HOST)
up-migration: DBI ?= outbound-us
up-migration: DB ?= wp-dev
up-migration:
	@./main.sh up-migration $(CFG) $(PRJ) $(DBH) $(DBI) $(DB)

# Preforms 1 DB migration rollback (down).
.PHONY: down-migration
down-migration: CFG ?= wp-spanner-emulator
down-migration: PRJ ?= test-project
down-migration: DBH ?= $(SPANNER_EMULATOR_HOST)
down-migration: DBI ?= outbound-us
down-migration: DB ?= wp-dev
down-migration:
	@./main.sh down-migration $(CFG) $(PRJ) $(DBH) $(DBI) $(DB)

# Run GCP emulators.
.PHONY: run-local-env
run-local-env:
	@docker-compose up -d

# Stop GCP emulators.
.PHONY: stop-local-env
stop-local-env:
	docker-compose down

# Generate mocks
.PHONY: generate-mock
generate-mock:
	@echo 'Generating mocks...'
	go generate -v ./...
