Cluster
-

````
export host=localhost
export port=9200
export index=megacorp
export type=employee
````

Running Elasticsearch:
````
./bin/elasticsearch

sudo /etc/init.d/elasticsearch status
sudo /etc/init.d/elasticsearch restart

curl 'http://$host:$port/?pretty'
````
````
# enable scripting
# appent into file /etc/elasticsearch/elasticsearch.yml
````

Shut down
````
curl -XPOST 'http://$host:$port/_shutdown'
````

#### [Upgrade](https://www.elastic.co/guide/en/elasticsearch/reference/current/setup-upgrade.html).

````json
curl $host:$port'/_cat/health?v'
curl $host:$port'/_cat/nodes?v'

# show shards
curl $host:$port/_cat/shards?v

# master
curl http://$host:$port/_cat/master?v

# node
curl http://$host:$port/_cat/nodeattrs?v

# nodes
curl http://$host:$port/_cat/nodes?v

# DISK SPACE available in your cluster 💿 .
curl -s $host:$port'/_cat/allocation?v'

# STATS
curl $host:$port/$index/'_stats?pretty'
curl $host:$port'/_nodes/stats/process?pretty'

# Local
curl $host:$port/_nodes/_local?pretty
curl $host:$port/_cluster/health?pretty

curl -XGET http://$host:$port/_cluster/stats?pretty

curl -XGET http://$host:$port/_nodes?pretty

curl -XGET http://$host:$port/_nodes/stats?pretty

# Cluster Settings
curl -XGET $host:$port/_cluster/settings

# Cluster Update Settings
curl -XPUT $host:$port/_cluster/settings -d '{
    "persistent" : {
        "discovery.zen.minimum_master_nodes" : 2
    }
}'

# tasks
curl -XGET http://$host:$port/_tasks?pretty
````

````json
# create index
curl -XPUT http://$host:$port/$index/

# delete index
curl -XDELETE http://$host:$port/$index/

# get index settings
curl $host:$port/$index/_settings?pretty
# or
# BEST ONE + SIZES !!! 👍
curl http://$host:$port/_cat/indices?v

# get indexes
curl http://$host:$port/_cat/indices?v

# get all mappings (types)
curl -XGET http://$host:$port/_mapping?pretty

# get mapping
curl -XGET http://$host:$port/$index/_mapping/$type

# put mapping for employee
curl -XPUT http://$host:$port/$index/_mapping/$type -d '{
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

# put mapping from file
curl -XPUT http://localhost:9200/ziipr -d @/vagrant/vagrant/elasticsearch.mapping.json

# IMPORTANT! Fields in the same index with the same name in two different types must have the same mapping
# Next code will spawn error
curl -XPUT http://$host:$port/test/ -d '{
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
# curl -XDELETE 'http://$host:$port/$index/$type'

# create alias
curl -XPOST $host:$port/_aliases -d '{
"actions": [
    {"add": {"alias": "megacorp", "index": "megacorp_v1"}}
]
}'

# get aliases
curl -XGET $host:$port/_alias/
curl -XGET $host:$port/_cat/aliases?v
````
