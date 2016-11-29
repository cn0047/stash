Cluster
-

Running Elasticsearch:
````
./bin/elasticsearch

sudo /etc/init.d/elasticsearch status
sudo /etc/init.d/elasticsearch restart

curl 'http://localhost:9200/?pretty'
````
````
# enable scripting
# appent into file /etc/elasticsearch/elasticsearch.yml
script.engine.groovy.inline.search: on
script.engine.groovy.inline.update: on
````

Shut down
````
curl -XPOST 'http://localhost:9200/_shutdown'
````

#### [Upgrade](https://www.elastic.co/guide/en/elasticsearch/reference/current/setup-upgrade.html).

````json
curl 'localhost:9200/_cat/health?v'
curl 'localhost:9200/_cat/nodes?v'
curl 'localhost:9200/_nodes/stats/process?pretty'

# show shards
curl localhost:9200/_cat/shards?v

# create index
curl -XPUT http://localhost:9200/megacorp/

# delete index
curl -XDELETE http://localhost:9200/ziipt/

# get index settings
curl localhost:9200/megacorp/_settings
# or
curl localhost:9200/_cat/indices?v

# master
curl http://localhost:9200/_cat/master?v
# node
curl http://localhost:9200/_cat/nodeattrs?v
# nodes
curl http://localhost:9200/_cat/nodes?v

# get indexes
curl http://localhost:9200/_cat/indices?v

# get all mappings (types)
curl -XGET http://localhost:9200/_mapping?pretty

# get mapping
curl -XGET http://localhost:9200/megacorp/_mapping/employee

# put mapping for employee
curl -XPUT http://localhost:9200/megacorp/_mapping/employee -d '{
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

# IMPORTANT! Fields in the same index with the same name in two different types must have the same mapping
# Next code will spawn error
curl -XPUT http://localhost:9200/test/ -d '{
"mappings" : {
      "boxing": {
        "properties": {"email": {"type": "string", "index": "not_analyzed"}}
      },
      "footbal": {
        "properties": {"email": {"type": "nested"}}
      }
}
}'

# delete mapping, from v2.3 NOT possible, need delete whole index
# curl -XDELETE 'http://localhost:9200/megacorp/employee'

# create alias
curl -XPOST localhost:9200/_aliases -d '{
"actions": [
    {"add": {"alias": "megacorp", "index": "megacorp_v1"}}
]
}'
````

````
# Local
curl localhost:9200/_nodes/_local?pretty
curl localhost:9200/_cluster/health?pretty

curl -XGET http://localhost:9200/_cluster/stats

curl -XGET 'http://localhost:9200/_nodes'

curl -XGET 'http://localhost:9200/_nodes/stats?pretty'

# Cluster Settings
curl -XGET localhost:9200/_cluster/settings

# Cluster Update Settings
curl -XPUT localhost:9200/_cluster/settings -d '{
    "persistent" : {
        "discovery.zen.minimum_master_nodes" : 2
    }
}'

# tasks
curl -XGET 'http://localhost:9200/_tasks?pretty'

Disk space available in your cluster#
curl -s 'localhost:9200/_cat/allocation?v'
````
