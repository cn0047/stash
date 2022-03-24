Reindex
-

#### To put document into index use:

````sh
# v7
PUT  /$idx/_doc/$id
POST /$idx/_doc/
PUT  /$idx/_create/$id
POST /$idx/_create/$id
````

#### To reindex all documents:

````sh
jh='Content-Type: application/json'

h='localhost:9200'
idx=megacorp
newindex=megacorp2
alias=megacorpalias
type=employee



# 1: Check indexes.
curl $h/_cat/indices?v

# 2: Put mapping for new index.
curl -XPUT $h/$newidx/ -d '{
  "mappings" : {}
}'

# 3: Update alias to point to new index.
curl -XPOST -H $jh $h/_aliases -d '{
"actions": [
    { "remove": { "index": "'$idx'", "alias": "'$alias'" }},
    { "add":    { "index": "'$newidx'", "alias": "'$alias'" }}
]
}'

# 4: Check alias.
curl -XGET $h/_cat/aliases?v
````
