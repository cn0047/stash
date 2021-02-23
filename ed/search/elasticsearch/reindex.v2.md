Reindex (with ALIAS) wihout downtime
-

## Reindex (with ALIAS) wihout downtime [APPROACH 1]

````sh
export h='localhost:9200'
export idx=megacorp
export alias=megacorp2
export type=employee
````
````sh
curl $h/_cat/indices?v

curl -XPUT $h/$alias/ -d '{
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

curl -XGET $h/$idx/$type/_count

# Run script which is re-index data into ES.

# Compare this value with returned form $idx.
curl -XGET $h/$alias/$type/_count

curl -XPOST $h/_aliases -d '{
"actions": [
    {"add": {"alias": "'$idx'", "index": "'$alias'"}}
]
}'

curl -XDELETE $h/$idx/

curl -XGET $h/_alias/
curl -XGET $h/_cat/aliases?v
````

## Reindex (with ALIAS) wihout downtime [APPROACH 2]

````sh
export h='localhost:9200'
export idx=megacorp
export newindex=megacorp2
export alias=megacorpalias
export type=employee
````
````sh
# 1: Check indexes.
curl $h/_cat/indices?v

# 2: Put mapping.
curl -XPUT $h/$newidx/ -d '{
  "mappings" : {}
}'

# 3: Update alias to point to new index.
curl -XPOST $h/_aliases -d '{
"actions": [
    { "remove": { "index": "'$idx'", "alias": "'$alias'" }},
    { "add":    { "index": "'$newidx'", "alias": "'$alias'" }}
]
}'

# 4: Check alias.
curl -XGET $h/_cat/aliases?v
````
