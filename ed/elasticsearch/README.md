Elasticsearch
-
2.2
1.6.0

Running Elasticsearch
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

In Elasticsearch, a document belongs to a type, and those types live inside an index.
You can draw some (rough) parallels to a traditional relational database:
````
Relational DB ⇒ Databases ⇒ Tables ⇒ Rows      ⇒ Columns
Elasticsearch ⇒ Indices   ⇒ Types  ⇒ Documents ⇒ Fields
````
We could use the DELETE verb to delete the document.
And the HEAD verb to check whether the document exists.
To replace an existing document - just PUT it again.

Use `delete-by-query` plugin to delete all documents matching a specific query.

`_score` field in the search results - the higher the score,
the more relevant the document is, the lower the score, the less relevant the document is.
`filter` clauses which allow to use a query
to restrict the documents that will be matched by other clauses, without changing how scores are computed.

[Upgrade](https://www.elastic.co/guide/en/elasticsearch/reference/current/setup-upgrade.html).

````json
curl 'localhost:9200/_cat/health?v'
curl 'localhost:9200/_cat/nodes?v'
curl 'localhost:9200/_nodes/stats/process?pretty'

# create index
curl -XPUT http://localhost:9200/megacorp/

# delete index
curl -XDELETE http://localhost:9200/megacorp/

# get indexes
curl http://localhost:9200/_cat/indices?v

# get all mappings (types)
curl -XGET http://localhost:9200/_mapping

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

# delete mapping
curl -XDELETE 'http://localhost:9200/megacorp/employee'

# create alias
curl -XPOST localhost:9200/_aliases -d '{
"actions": [
    {"add": {"alias": "megacorp", "index": "megacorp_v1"}}
]
}'
````

https://www.elastic.co/guide/en/elasticsearch/reference/current/date-math-index-names.html