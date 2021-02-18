Administration
-

````sh
host=localhost
port=9200
index=
ctj='Content-Type: application/json'
fm=/tmp/mapping.json

# create index
curl -XPUT $host:$port/$index | jq

# put mapping
curl -XPUT $host:$port/$index/_mapping -H $ctj -d @$fm | jq

# delete index
curl -XDELETE $host:$port/$index | jq

# get mapping
curl -XGET $host:$port/$index/_mapping | jq

curl -XGET $host:$port/$index/_count | jq
````

#### Reindex

````sh
index=metrics
indexNew=metrics_20200131

# create new index
curl -XPUT $host:$port/$indexNew | jq
# put mapping into new index
curl -XPUT $host:$port/$indexNew/_mapping -H $ctj -d @$fm | jq
# reindex
curl -XPOST $host:$port/_reindex -H $ctj -d '{
  "source": {"index": "'$index'"},
  "dest": {"index": "'$indexNew'"}
}' | jq
# check counts
curl -XPOST "$host:$port/$index/_count" | jq
curl -XPOST "$host:$port/$indexNew/_count" | jq
# alias
curl -XPOST $host:$port/_aliases -H $ctj -d '{
  "actions": [{"add": {"alias": "'$index'", "index": "'$indexNew'"}}]
}' | jq

````
