#!/usr/bin/env bash

commandExists() {
    if [ -x "$(command -v $1)" ]; then
      echo 1
    else
      echo 0
    fi
}

result=$(commandExists $1)
echo "$result"
