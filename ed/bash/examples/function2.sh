#!/bin/bash

hello () {
    echo "Hello $1"
    return
}

result=$(hello "world")
echo "$result"
