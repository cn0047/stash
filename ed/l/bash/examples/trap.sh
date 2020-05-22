#!/bin/bash

function finally() {
  echo -e "\nFinally."
}

function cleanup() {
  echo -e "\nClean up."
}

trap cleanup INT TERM QUIT # call func cleanup if script is stopped
                           # ⚠️ only linux
trap finally EXIT          # @TODO: *.

for i in $(seq 1 10); do echo -n . ; sleep 0.1; done
# while true;           do echo -n . ; sleep 0.1; done
