Aggregations (Analytics)
-

````sh
curl -XPOST "$host:$port/$index/_search?size=0" -H $ctj -d '{
    "aggs" : {
        "aggregated" : {
            "filter" : {
              "bool": {"must": [
                {"match": {"name": "'$m'"}},
                {"term": {"properties.by_user_id": "'$uId'"}},
                {"range" : {"timestamp": {"from": 1178323112, "to": 2579519655}}}
              ]}
            },
            "aggs": {
                "aggregated": {"avg": {"field" : "properties.value"}}
            }
        }
    }
}' | jq
````
