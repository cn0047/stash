Elasticsearch
-
2.3.3
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

# show shards
curl localhost:9200/_cat/shards?v

# create index
curl -XPUT http://localhost:9200/megacorp/

# delete index
curl -XDELETE http://localhost:9200/megacorp/

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

# Fields in the same index with the same name in two different types must have the same mapping
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

# delete mapping
curl -XDELETE 'http://localhost:9200/megacorp/employee'

# create alias
curl -XPOST localhost:9200/_aliases -d '{
"actions": [
    {"add": {"alias": "megacorp", "index": "megacorp_v1"}}
]
}'
````

#### Versioning

Internally, Elasticsearch has marked the old document as deleted and added an entirely new document.
The old version of the document doesn’t disappear immediately,
although you won’t be able to access it.
Elasticsearch cleans up deleted documents in the background as you continue to index more data.


[Meta-Fields](https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-fields.html).

[Groovy](https://www.elastic.co/guide/en/elasticsearch/reference/current/modules-scripting.html#_document_fields).

[Glossary](https://www.elastic.co/guide/en/elasticsearch/reference/current/glossary.html).

A shard is a single Lucene instance.
An index is a logical namespace which points to primary and replica shards.
