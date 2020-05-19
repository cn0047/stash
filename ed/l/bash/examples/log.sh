#!/bin/bash

curl -s -XPOST 'https://realtimelog.herokuapp.com:443/r20qpxrrto' \
  -H 'Content-Type: application/json' -d '{"msg": "'$1'"}'
