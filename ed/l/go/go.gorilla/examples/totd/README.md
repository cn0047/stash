TD
-

## How to develop

````sh
# prepare local env
make spanner-start
make spanner-load-fixtures # load predefined data into DB

# run on local env
export SPANNER_EMULATOR_HOST=localhost:9010
go run ./cmd/poc-td/main.go -config-path=config/local.json

# check
curl -s 'http://localhost:9000/health'

# build go project
make build

# run unit tests
make test-unit

# prepare env and run BDD tests
make test-spec

# run BDD tests (only run tests, without env preparing)
make test-spec-run
````
