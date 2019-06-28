Megacorp examples
-

````sh
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

````sh
curl -XPUT 'http://localhost:9200/twitter/tweet/1?ttl=1m' -d '{
    "user": "kimchy",
    "message": "Trying out elasticsearch, so far so good?"
}'
````
