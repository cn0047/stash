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
````
# enable scripting
# appent into file /etc/elasticsearch/elasticsearch.yml
script.engine.groovy.inline.search: on
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

````json
# create index
curl -XPUT http://localhost:9200/megacorp/

# delete index
curl -XDELETE http://localhost:9200/megacorp/

# get indexes
curl http://localhost:9200/_cat/indices?v

# get mapping
curl -XGET http://localhost:9200/ziipr/_mapping/users

# put mapping for user
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
````

````json
// Create new documents (employee)
curl -XPUT localhost:9200/megacorp/employee/1 -d '{
"first_name" : "John",
"last_name" : "Smith",
"age" : 25,
"about" : "I love to go rock climbing",
"last_login_at": "2016-01-21",
"city": "London",
"location": {"lat": 51.5072, "lon": 0.1275},
"interests": [ "sports", "music" ]
}'
curl -XPUT localhost:9200/megacorp/employee/2 -d '{
"first_name" : "Jane",
"last_name" : "Smith",
"age" : 32,
"about" : "I like to collect rock albums",
"last_login_at": "2016-01-10",
"city": "Manchester",
"location": {"lat": 53.4667, "lon": 2.2333},
"interests": [ "music" ]
}'
curl -XPUT localhost:9200/megacorp/employee/3 -d '{
"first_name" : "Douglas",
"last_name" : "Fir",
"age" : 35,
"about": "I like to build cabinets",
"last_login_at": "2016-02-14",
"city": "Kyiv",
"location": {"lat": 50.4500, "lon": 30.5233},
"interests": [ "forestry" ]
}'
curl -XPUT localhost:9200/megacorp/employee/4 -d '{
"first_name" : "Louis",
"last_name" : "de Funès",
"age" : 70,
"about": "actor",
"last_login_at": "2012-03-04",
"city": "Paris",
"location": {"lat": 48.8567, "lon": 2.3508},
"interests": [ "fantomas", "theatre" ]
}'
curl -XPUT localhost:9200/megacorp/employee/5 -d '{
"first_name" : "Cristiano",
"last_name" : "Ronaldo",
"age" : 31,
"about": "footballer",
"last_login_at": "2016-03-11",
"city": "Santo António",
"location": {"lat": 37.1939, "lon": 7.4158},
"interests": [ "football", "cars", "casino" ]
}'
curl -XPUT localhost:9200/megacorp/employee/6 -d '{
"first_name" : "Gennady",
"last_name" : "Golovkin",
"age" : 33,
"about": "Professional Boxer",
"last_login_at": "2016-03-03",
"city": "Karaganda",
"location": {"lat": 49.8333, "lon": 73.1667},
"interests": [ "boxing", "WBA", "IBO", "cars" ]
}'
curl -XPUT localhost:9200/megacorp/employee/7 -d '{
"first_name" : "Jackie",
"last_name" : "Chan",
"age" : 61,
"about": "Martial Artist",
"last_login_at": "2016-03-12",
"city": "Hong Kong",
"location": {"lat": 22.2783, "lon": 114.1747},
"interests": [ "movie", "hollywood", "kong foo" ]
}'
// Get employee 1
curl -XGET localhost:9200/megacorp/employee/1
````

We could use the DELETE verb to delete the document.
And the HEAD verb to check whether the document exists.
To replace an existing document - just PUT it again.

#### Search

````json
# find all employee
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
        "bool": {
            "filter" : [
                {"range" : {"age": {"gt": 32}}}
            ]
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

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {"bool": {
        "filter": {
            "script": {
                "script": "doc['"'"'age'"'"'].value > 33"
            }
        }
    }}
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query" : {"bool": {
        "filter": {
            "script": {
                "script": {
                    "inline": "doc['"'"'age'"'"'].value > param1",
                    "params" : {"param1": 40}
                }
            }
        }
    }}
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "fields" : ["city"],
    "query" : {
        "bool": {
            "filter": {
                "script": {
                    "script": {
                        "inline": "doc['"'"'location'"'"'].distanceInKm(lat, lon)",
                        "params": {"lat": 49.8333, "lon": 73.1667}
                    }
                }
            }
        }
    }
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "sort" : [{ "city" : "asc" }]
}'

# Sorting by distance from London.
curl -XGET localhost:9200/megacorp/employee/_search -d '{
  "sort": [
    {
      "_geo_distance": {
        "location": {
          "lat":  51.5072,
          "lon": 0.1275
        },
        "order":         "asc",
        "unit":          "km",
        "distance_type": "plane"
      }
    }
  ]
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
#

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
