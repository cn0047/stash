#!/usr/bin/env bash

function prn() {
  local i=0
  for ((i=0; i<10; i++)); do
      echo -n '.'
  done
  printf "\n"
}

for ((i=0; i<10; i++)); do
    prn
done
