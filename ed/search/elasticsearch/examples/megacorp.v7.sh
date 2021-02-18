#!/bin/sh

host=localhost
port=9200
index=

# create index
curl -XPUT $host:$port/$index | jq
