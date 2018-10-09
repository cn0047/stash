#!/usr/bin/env bash

declare -ra arr=(
    "ed/go/examples/whatever/rand.2.go"
    "ed/go/examples/whatever/rand.2.go"
    "ed/go/examples/whatever/rand.2.go"
)

f1() {
  for i in "${arr[@]}"; do
      go run "$i" &
  done
  wait
}

f2() {
  rm -rf /tmp/bash.inParallel.out

  for i in ${arr[@]}; do
      go run ${i} 2>>/tmp/bash.inParallel.out &
  done
  wait

  # cat /tmp/bash.inParallel.out
  slowCount=$(grep IT_WAS_SLOW -c /tmp/bash.inParallel.out);
  echo "slowCount = ${slowCount}"
  if [ ${slowCount} -ne 0 ]; then
    echo "It was slow."
  fi
}

echo "Start:"
f2
echo "End."
