#!/bin/sh

host=localhost
port=9200
index=metrics

curl -XDELETE $host:$port/$index | jq

# create index
curl -XPUT $host:$port/$index | jq
# put mapping
curl -XPUT $host:$port/$index/_mapping -H 'Content-Type: application/json' -d '{
    "properties": {
        "msg": {"type": "text"},
        "tiq": {"type": "integer"}
    }
}' | jq

# str vals
curl -XPOST $host:$port/$index/_doc -H 'Content-Type: application/json' -d '{
    "msg": "test",
    "tiq" : "257"
}' | jq
curl -XPOST $host:$port/$index/_doc -H 'Content-Type: application/json' -d '{
    "msg": "test",
    "tiq" : "95"
}' | jq

curl -XPOST $host:$port/$index/_doc -H 'Content-Type: application/json' -d '{
    "msg": "test",
    "tiq" : "1o2"
}' | jq

# int vals
curl -XPOST $host:$port/$index/_doc -H 'Content-Type: application/json' -d '{
    "msg": "test",
    "tiq" : 25
}' | jq
curl -XPOST $host:$port/$index/_doc -H 'Content-Type: application/json' -d '{
    "msg": "test",
    "tiq" : 95
}' | jq



curl -XGET $host:$port/$index/_mapping | jq
curl -XGET $host:$port/$index/_count | jq
curl -XGET $host:$port/$index/_search | jq '.hits.hits'

curl -XPOST "$host:$port/$index/_search?size=0" -H 'Content-Type: application/json' -d '{
    "aggs" : {
        "max_tiq" : { "max" : { "field" : "tiq" } }
    }
}' | jq '.aggregations.max_tiq'

curl -XPOST "$host:$port/_sql?format=txt" -H 'Content-Type: application/json' -d '{
  "query": "select max(tiq) from '$index'"
}'
