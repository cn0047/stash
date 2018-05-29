#!/usr/bin/env bash

# Script to run commands related to tests.
#
# @example . test.sh
# @example . test.sh cover

declare -a unitTests=(
    "transactional/users/models"
    "transactional/users/service"
    "transactional/favorites/models"
    "transactional/favorites/service"
    "transactional/favorites.v2/models"
    "transactional/favorites.v2/service"
    "content/ratings/models"
    "content/ratings/service"
    "transactional/dynamic-config/model"
    "transactional/dynamic-config/service"
    "transactional/contact/models"
    "transactional/contact/service"
    "transactional/consent/service"
)

declare -a integrationTests=(
    "transactional/users/internal"
    "content/ratings/internal"
    "transactional/contact/internal"
)

# Function to initialize tests (bootstrapping).
setUp() {
    echo "Init tests"
    set -e
    source ./env.sh

    # Stop test server (with purpose to be on safe side).
    curl http://localhost:8222/quit

    # Init test server.
    go run src/common/api/apitest/cmd/main.go
}

# Function to run all tests (unit & integration).
runAllTests() {
    echo "Running unit tests"
    for i in "${unitTests[@]}"; do
       goapp test -cover "$i"
    done

    echo "Running integration tests"
    for i in "${integrationTests[@]}"; do
       goapp test -cover "$i"
    done
}

# Function to generate coverage report based on unit tests.
generateCoverageReport() {
    # Create temporary directory for internal coverage stuff.
    mkdir -p ./.cover

    # Generate coverage parts based on unit tests.
    for i in "${unitTests[@]}"; do
        part=${i////_}".part"
        goapp test -cover -coverprofile "./.cover/$part" "$i"
    done

    # Generate final coverage report.
    echo "mode: set" > ./.cover/coverage.out
    grep -h -v "mode: set" ./.cover/*.part >> ./.cover/coverage.out
    go tool cover -html=./.cover/coverage.out -o=coverage.html

    rm -rf ./.cover
}

# This block executes command
# appropriate to received "script execution argument".
# @see Line 5.
# @see Line 6.
case $1 in
# This command generates coverage report.
cover)
    setUp
    generateCoverageReport
    ;;
# This (default) command runs all test cases.
*)
    setUp
    runAllTests
    ;;
esac
