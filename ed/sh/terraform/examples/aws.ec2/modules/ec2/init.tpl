#!/bin/bash

curl -XPOST 'https://realtimelog.herokuapp.com:443/dyhnt08p53n' \
  -H 'Content-Type: application/json' -d '{
    "tf":"aws.x", "msg": "init", "dns": "'${dns}'","env":"'env'"
  }'
