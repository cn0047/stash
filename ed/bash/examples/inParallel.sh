#!/usr/bin/env bash

declare -a ar1=(
    "ed/go/examples/whatever/rand.2.go"
    "ed/go/examples/whatever/rand.2.go"
    "ed/go/examples/whatever/rand.2.go"
)

echo "Start:"

for i in "${ar1[@]}"; do
   go run "$i" &
done

wait

echo "End."
