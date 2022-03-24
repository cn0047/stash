Administration
-

States:
* green - all ok.
* yellow - some replicas shards are missing
  but primaries are all available and serving data.
* red - at least 1 primary and all it\'s replicas are offline,
  data posibly broken.

Node types (in ESv2.2):
* master - brain of the whole operatioln.
* data
* client - gateway to your cluster (node.master and node.data eq to false). Need half of data node capacity.
* tribe - special type of client node that can connect to multiple clusters.

cluster = 1 client + 3 master + 4 data

`export ES_HEAP_SIZE=32g` - for best performance
`bootstrap.mlockall: true` - no swap files

For OS:
````
File Descriptors: 64000
MMap: unlimited
````

Temporary disable shards rebalance:
````sh
host=localhost
port=9200
index=megacorp
type=employee

h=$host:$port
h=localhost:9200



jh='Content-Type: application/json'

curl -XPUT $h/_cluster/settings -d '{
    "transient" : {
        "cluster.routing.allocation.enable" : "none",
        "cluster.routing.allocation.enable" : "all"
    }
}'
````

Running Elasticsearch:
````sh
./bin/elasticsearch

sudo /etc/init.d/elasticsearch status
sudo /etc/init.d/elasticsearch restart

curl '$h/?pretty'
````
````sh
# enable scripting
# appent into file /etc/elasticsearch/elasticsearch.yml
script.groovy.sandbox.enabled: true
````

Shut down:
````sh
curl -XPOST '$h/_shutdown'
````

#### [Upgrade](https://www.elastic.co/guide/en/elasticsearch/reference/current/setup-upgrade.html).

````sh
curl "$h/_cat/health?v"

# show shards
curl "$h/_cat/shards?v"

# master
curl "$h/_cat/master?v"

# node
curl "$h/_cat/nodeattrs?v"

# nodes
curl "$h/_cat/nodes?v"
curl $h/_nodes | jq '.nodes| keys[]'

# DISK SPACE available in your cluster 💿 .
curl -s "$h/_cat/allocation?v"

# STATS
curl "$h/$idx/_stats?pretty" | jq
curl "$h/_nodes/stats/process?pretty" | jq

# Local
curl "$h/_nodes/_local?pretty"
curl "$h/_cluster/health?pretty"

curl "$h/_cluster/stats?pretty"

curl "$h/_nodes?pretty"

curl "$h/_nodes/stats?pretty"

# Cluster Settings
curl "$h/_cluster/settings" | jq

# Cluster Update Settings
# temporary
curl -XPUT "$h/_cluster/settings" -d '{
    "transient" : {
        "discovery.zen.minimum_master_nodes" : 2
    }
}'
# persistent
curl -XPUT "$h/_cluster/settings" -d '{
    "persistent" : {
        "discovery.zen.minimum_master_nodes" : 2
    }
}'

# tasks
curl "$h/_tasks?pretty"
````

````sh
# create index
curl -XPUT $h/$idx/

# delete index
curl -XDELETE $h/$idx/

# delete all documents from index
curl -XDELETE $h/$idx -H $jh -d '{
  "query": {"match_all": {}}
}'

# get index settings
curl "$h/$idx/_settings?pretty"

# get all indexes
# BEST ONE + SIZES !!! 👍
curl "$h/_cat/indices?v"

# get all mappings (types)
curl $h/_mapping

# get mapping v2
curl $h/$idx/_mapping/$type
curl $h/$idx/_mapping | jq
# get mapping v7
curl -XGET $h/$idx/_mapping | jq

curl -XGET $h/$idx/_count | jq

# put mapping for type employee
curl -XPUT $h/$idx/_mapping/$type -d '{
  "employee": {
      "properties": {
          "first_name": {"type": "string"},
          "last_name": {"type": "string"},
          "age": {"type": "integer"},
          "about": {"type": "string"},
          "last_login_at": {"type": "date", "format": "yyy-MM-dd"},
          "city": {"type": "string"},
          "location": {"type": "geo_point", "lat_lon": "true"},
          "interests": {"type": "string"}
      }
  }
}'

# put mapping from file v2
curl -XPUT $h/zii -H $jh -d @/vagrant/vagrant/elasticsearch.mapping.json
# put mapping from file v7
m=/tmp/mapping.json
curl -XPUT $h/$inx/_mapping -H $jh -d @$m | jq

# IMPORTANT: Fields in the same index with the same name in two different types must have the same mapping
# Next code will spawn error
curl -XPUT $h/test/ -d '{
"mappings" : {
      "boxing": {
        "properties": {"email": {"type": "string", "index": "not_analyzed"}}
      },
      "footbal": {
        "properties": {"email": {"type": "nested"}}
      }
}
}'

# IMPORTANT: delete mapping, from v2.3 NOT possible, need delete whole index
# curl -XDELETE '$h/$idx/$type'

# create alias
curl -XPOST $h/_aliases -d '{
"actions": [
    {"add": {"alias": "megacorp", "index": "megacorp_v1"}}
]
}'

# update alias
curl -XPUT $h/$idx/_alias/$a

# delete alias
curl -XPUT $h/$idx/_alias/$a

# get aliases
curl "$h/_alias/" | jq
curl "$h/_cat/aliases?v"
````

````sh
# Reindex whole index !!!
curl -XPOST $h/_reindex -d '{
  "source": {"index": "megacorp"},
  "dest": {"index": "megacorp_2"}
}'

curl -XPOST $h/_reindex -d '{
  "source": {"index": "megacorp", "type": "employee"},
  "dest": {"index": "new_megacorp", "type": "new_employee"}
}'
````

````sh
# Enabling caching per request
curl "$h/$idx/_search?request_cache=true" -d'
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
curl "$h/_stats/request_cache?pretty&human"
# or
curl "$h/_nodes/stats/indices/request_cache?pretty&human"

# Clear Cache
curl -XPOST "$h/$idx/_cache/clear"
# or
curl -XPOST "$h/kimchy,elasticsearch/_cache/clear?request_cache=true"

# Flush
curl -XPOST "$h/$idx/_flush"

# Refresh index
curl -XPOST "$h/$idx/_refresh"
````
