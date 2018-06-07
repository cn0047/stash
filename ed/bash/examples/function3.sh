#!/usr/bin/env bash

# Return from function fast.
f() {
  if [ $1 = '0' ]; then
    return
  fi
  echo $1

}

f "1"
