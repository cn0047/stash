Reindex (with ALIAS) wihout downtime
-

export host='localhost'
export port=9201
export index=indexName
export alias=aliasName
export type=typeName

## Reindex (with ALIAS) wihout downtime

````json
curl http://$host:$port/_cat/indices?v

curl -XPUT http://$host:$port/$alias/ -d '{
  "mappings" : {
    "'$typeName'": {
      "properties": {
      }
    }
  }
}'

curl -XGET $host:$port/$index/$type/_count

# Run script which is re-index data into ES.

# Compare this value with returned form $index.
curl -XGET $host:$port/$alias/$type/_count

curl -XDELETE http://$host:$port/$index/

curl -XPOST localhost:$port/_aliases -d '{
"actions": [
    {"add": {"alias": "'$index'", "index": "'$alias'"}}
]
}'

curl -XGET $host:$port/_alias/
curl -XGET $host:$port/_cat/aliases?v
````
