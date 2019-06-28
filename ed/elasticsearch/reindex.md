Reindex (with ALIAS) wihout downtime
-

## Reindex (with ALIAS) wihout downtime [APPROACH 1]

````sh
export host='localhost'
export port=9200
export index=megacorp
export alias=megacorp2
export type=employee
````
````sh
curl $host:$port/_cat/indices?v

curl -XPUT $host:$port/$alias/ -d '{
  "mappings" : {
    "'$type'": {
      "properties": {
        "first_name": {"type": "string", "index": "not_analyzed"},
        "last_name": {"type": "string", "index": "not_analyzed"},
        "age": {"type": "integer"},
        "about": {"type": "string", "index": "not_analyzed"},
        "last_login_at": {"type": "date", "format": "yyy-MM-dd"},
        "city": {"type": "string", "index": "not_analyzed"},
        "location": {"type": "geo_point"},
        "interests": {"type": "string"},
        "fetish": {"type": "nested"},
        "pictures": {"type": "nested"}
      }
    }
  }
}'

curl -XGET $host:$port/$index/$type/_count

# Run script which is re-index data into ES.

# Compare this value with returned form $index.
curl -XGET $host:$port/$alias/$type/_count

curl -XDELETE $host:$port/$index/

curl -XPOST $host:$port/_aliases -d '{
"actions": [
    {"add": {"alias": "'$index'", "index": "'$alias'"}}
]
}'

curl -XGET $host:$port/_alias/
curl -XGET $host:$port/_cat/aliases?v
````

## Reindex (with ALIAS) wihout downtime [APPROACH 2]

````sh
export host='localhost'
export port=9200
export index=megacorp
export newindex=megacorp2
export alias=megacorpalias
export type=employee
````
````sh
curl $host:$port/_cat/indices?v

# Put mapping

curl -XPUT $host:$port/$index/_alias/$alias

curl -XGET $host:$port/_cat/aliases?v

# Put mapping into newindex

curl -XPOST $host:$port/_aliases -d '{
"actions": [
    { "remove": { "index": "'$index'", "alias": "'$alias'" }},
    { "add":    { "index": "'$newindex'", "alias": "'$alias'" }}
]
}'
````
