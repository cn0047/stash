Elasticsearch
-
1.6.0

elasticsearch.org/download
````
curl -L -O http://download.elasticsearch.org/PATH/TO/VERSION.zip
unzip elasticsearch-$VERSION.zip
cd elasticsearch-$VERSION
````

Installation
````
wget https://download.elasticsearch.org/elasticsearch/elasticsearch/elasticsearch-0.90.7.deb
sudo dpkg -i elasticsearch-0.90.7.deb
````

Installing Marvel
````
./bin/plugin -i elasticsearch/marvel/latest

http://localhost:9200/_plugin/marvel/
http://localhost:9200/_plugin/marvel/sense/
````

Running Elasticsearch
````
./bin/elasticsearch

sudo /etc/init.d/elasticsearch status
sudo /etc/init.d/elasticsearch restart

curl 'http://localhost:9200/?pretty'
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

````

# create index
curl -XPUT http://localhost:9200/megacorp/

# delete index
curl -XDELETE http://localhost:9200/megacorp/

# get indexes
curl http://localhost:9200/_cat/indices?v

# get mapping
curl -XGET http://localhost:9200/ziipr/_mapping/users

# put mapping for user
curl -XPUT http://localhost:9200/megacorp/_mapping/users -d '{
  "users": {
      "_id" : {"path" : "user_id"},
      "properties": {
          "user_id": {"type": "long"},
          "created_at": {"type": "date", "format": "yyy-MM-dd HH:mm:ss"}
      }
  }
}'
````

````json
// Create new document (employee id 1)
curl -XPUT localhost:9200/megacorp/employee/1 -d '{
"first_name" : "John",
"last_name" : "Smith",
"age" : 25,
"about" : "I love to go rock climbing",
"interests": [ "sports", "music" ]
}'

// Create new document (employee id 2)
curl -XPUT localhost:9200/megacorp/employee/2 -d '{
"first_name" : "Jane",
"last_name" : "Smith",
"age" : 32,
"about" : "I like to collect rock albums",
"interests": [ "music" ]
}'

// Create new document (employee id 3)
curl -XPUT localhost:9200/megacorp/employee/3 -d '{
"first_name" : "Douglas",
"last_name" : "Fir",
"age" : 35,
"about": "I like to build cabinets",
"interests": [ "forestry" ]
}'

````
// Get employee 1
curl -XGET localhost:9200/megacorp/employee/1
````

We could use the DELETE verb to delete the document.
And the HEAD verb to check whether the document exists.
To replace an existing document - just PUT it again.

#### Search

````json
curl -XGET localhost:9200/megacorp/employee/_search

curl -XGET localhost:9200/megacorp/employee/_search?q=last_name:Smith

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "filtered" : {
            "filter" : { "range" : { "age" : { "gt" : 30 } } }
        }
    }
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "filtered" : {
            "filter" : { "range" : { "age" : { "gt" : 30 } } },
            "query" : { "match" : { "last_name" : "smith" } }
        }
    }
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "filtered" : {
            "query" : { "match" : { "first_name" : "Douglas" } }
        }
    }
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "filtered" : {
            "query" : { "match_all" : {  } },
            "filter": { "term": { "age": 35 } }
        }
    }
}'
````

#### Full-text search

````json
curl -XGET localhost:9200/megacorp/employee/_search -d '{
"query" : { "match" : { "about" : "rock climbing" } }
}'
# we'll receive: "I love to go rock climbing" and "I like to collect rock albums"

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "match_phrase" : { "about" : "rock climbing" }
    }
}'

# we'll receive: "I love to go rock climbing"

curl -XGET localhost:9200/megacorp/employee/_search -d '
{
    "query" : {
        "match_phrase" : { "about" : "rock climbing" }
    },
    "highlight": { "fields" : { "about" : {} } }
}
'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {
        "wildcard" : { "about" : "*limbing" }
    }
}'

# we'll receive: "I love to go rock climbing"
````

#### Analytics

````json
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs": {
        "all_interests": { "terms": { "field": "interests" } }
    }
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query": {
        "match": { "last_name": "smith" }
    },
    "aggs": {
        "all_interests": {
            "terms": { "field": "interests" }
        }
    }
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "all_interests" : {
            "terms" : { "field" : "interests" },
            "aggs" : {
                "avg_age" : { "avg" : { "field" : "age" } }
            }
        }
    }
}'
````
