cluster
-

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
````
