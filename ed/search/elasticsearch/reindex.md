Reindex
-

#### To put document into index use:

````sh
PUT  /$idx/_doc/$id
POST /$idx/_doc/
PUT  /$idx/_create/$id
POST /$idx/_create/$id
````

#### To reindex all documents:

````sh
export h='localhost:9200'
export idx=megacorp
export newindex=megacorp2
export alias=megacorpalias
export type=employee



# 1: Check indexes.
curl $h/_cat/indices?v

# 2: Put mapping for new index.
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
