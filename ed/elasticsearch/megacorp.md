Megacorp examples
-

````json
# Create new documents (car)
curl -XPUT localhost:9200/megacorp/car/15?parent=15 -d '{
    "name" : "RANGE_ROVER_SPORT",
    "brand" : "land rover",
    "about": "cool car!"
}'

# Bulk insert (the possible actions are index, create, delete and update)
curl -XPOST 'localhost:9200/megacorp/employee/_bulk?pretty' -d '
{"index":{"_id":"8"}}
{"name": "John Doe" }
{"index":{"_id":"9"}}
{"name": "Jane Doe 2"}
'

# Bulk insert from file
curl -XPOST 'localhost:9200/megacorp/employee/_bulk?pretty' --data-binary "@/vagrant/megacorpEmployee.json"
````

````
# Enabling caching per request
curl 'localhost:9200/my_index/_search?request_cache=true' -d'
{
  "size": 0,
  "aggs": {
    "popular_colors": {
      "terms": {
        "field": "colors"
      }
    }
  }
}
'

# Monitoring cache usage
curl 'localhost:9200/_stats/request_cache?pretty&human'
# or
curl 'localhost:9200/_nodes/stats/indices/request_cache?pretty&human'

# Clear Cache
curl -XPOST 'http://localhost:9200/megacorp/_cache/clear'
# or
curl -XPOST 'localhost:9200/kimchy,elasticsearch/_cache/clear?request_cache=true'

# Flush
curl -XPOST 'http://localhost:9200/megacorp/_flush'

# Refresh index
curl -XPOST 'http://localhost:9200/megacorp/_refresh'
````

````json
# Reindex whole index
curl -XPOST localhost:9200/_reindex -d '{
  "source": {"index": "megacorp"},
  "dest": {"index": "megacorp_2"}
}'

curl -XPOST localhost:9200/_reindex -d '{
  "source": {"index": "megacorp", "type": "employee"},
  "dest": {"index": "new_megacorp", "type": "new_employee"}
}'
````

````json
curl -XPUT 'http://localhost:9200/twitter/tweet/1?ttl=1m' -d '{
    "user": "kimchy",
    "message": "Trying out elasticsearch, so far so good?"
}'
````
