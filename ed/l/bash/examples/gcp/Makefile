APP_NAME := $(shell basename "$(CURDIR)")

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



# v2

# Init migration tool.
.PHONY: migrate-init
migrate-init: OS ?= linux
migrate-init: ARCH ?= amd64
migrate-init:
	@ops/sh/spanner.sh migrate-init $(OS) $(ARCH)

# Start Spanner.
.PHONY: spanner-start
spanner-start: CFG ?= to
spanner-start: PRJ ?= test-project
spanner-start: DBI ?= test-instance
spanner-start: DB ?= test-db
spanner-start:
	@ops/sh/spanner.sh spanner-start $(CFG) $(PRJ) $(DBI) $(DB)

# Init Spanner.
.PHONY: spanner-init
spanner-init: CFG ?= to
spanner-init: PRJ ?= test-project
spanner-init: DBI ?= test-instance
spanner-init: DB ?= test-db
spanner-init:
	@ops/sh/spanner.sh spanner-init $(CFG) $(PRJ) $(DBI) $(DB)

# Stop Spanner.
.PHONY: spanner-stop
spanner-stop:
	@ops/sh/spanner.sh spanner-stop
