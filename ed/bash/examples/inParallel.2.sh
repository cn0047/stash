#!/usr/bin/env bash

declare -a ar1=(
    "ed/go/examples/whatever/rand.2.go"
    "ed/go/examples/whatever/rand.2.go"
    "ed/go/examples/whatever/rand.2.go"
)
pids=""

echo "Start:"

for i in "${ar1[@]}"; do
    go run "$i" &
    pids+=" $!"
done

for p in ${pids}; do
    if ! wait ${p}; then
        echo "PID $p FAILED"
        exit -1
    fi
done

echo "End."
