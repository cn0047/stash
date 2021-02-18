Search
-

````sh
h=localhost:9200
idx=

# get total count
curl -XGET $h/$idx/_count | jq

# find all
curl -XGET $h/$idx/_search | jq '.hits.hits'
````
