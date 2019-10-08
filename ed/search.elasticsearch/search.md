Search
-

[Dates math](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-range-query.html#_date_math_and_rounding).
[Regexp syntax](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-regexp-query.html#regexp-syntax).
[scroll - Pagination of continuously updating data](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-request-scroll.html).

General search types:
* Full-text: phrase, word, wildcard.
* Geo.
* Faceted (something like groups or categories).

Query context - How well does document match query (score).
Filter context - does document match query clause (yes/no).

The parameters allowed in the URI search
[are](https://www.elastic.co/guide/en/elasticsearch/reference/current/search-uri-request.html#_parameters_4)

There are three types of `match` query: `boolean`, `phrase`, and `phrase_prefix`.

Boolean:

* must - clause must appear in matching documents.
* should - may appear but may ney not.
* must_not
* filter - must appear in result but result not scored.

````sh
# Get employee 1
curl -XGET localhost:9200/megacorp/employee/1

# Multi get
curl 'localhost:9200/_mget?pretty' -d '{
    "docs" : [
        {"_index" : "megacorp", "_type" : "employee", "_id" : "1"},
        {"_index" : "megacorp", "_type" : "employee", "_id" : "2"},
        {"_index" : "megacorp", "_type" : "employee", "_id" : "3"}
    ]
}'

# Multi get by certain index and type
curl 'localhost:9200/megacorp/employee/_mget?pretty' -d '{
    "docs" : [{"_id" : "1"}, {"_id" : "2"}]
}'
# or
curl 'localhost:9200/megacorp/employee/_mget?pretty' -d '{
    "ids" : ["1", "2"]
}'
````

````sh
# find all employee
curl -XGET localhost:9200/megacorp/employee/_search

curl -XGET localhost:9200/megacorp/employee/_search?q=last_name:Smith

# validate query
curl -XGET localhost:9200/megacorp/employee/_validate/query -d '{
    "query": {"match_all" : {}}
}'

# calculate count of all documents
curl -XGET localhost:9200/megacorp/employee/_count -d '{
    "query": {"match_all" : {}}
}'

# explain
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "explain": true,
    "query": {"match_all" : {}}
}'
# or
curl -XGET localhost:9200/megacorp/employee/4/_explain?q=first_name:Louis&pretty

# profile
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "profile": true,
    "query": {"match_all" : {}}
}'

# version for each search hit
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "version": true,
    "query": {"match_all" : {}}
}'

# Search by AND condition
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "query": {
    "bool": {
      "must": [
        { "match": { "about": "like" } } ,
        { "match": { "city": "Kyiv" } }
      ]
    }
  }
}'

# Search all and filter result by AND condition
curl -XGET localhost:9200/megacorp/employee/_search?pretty -d '{
    "fields": ["last_login_at", "age"],
    "query" : {
        "bool": {
            "must": { "match_all": {} },
            "filter" : [
                {"range" : {"age": {"gt": 31}}},
                {"range" : {"last_login_at": {"gt": "2016-03-01"}}}
            ]
        }
    }
}'

# Array interests contains sports and music
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "query": {
    "bool": {
      "must": [
        { "term": { "interests": "sports" } } ,
        { "term": { "interests": "music" } }
      ]
    }
  }
}'
# or (using query_string)
curl -XGET localhost:9200/megacorp/employee/_search -d '{
"query" : {
    "query_string": {
      "query": "(interests:sports AND interests:music)"
    }
  }
}'

# Array interests contains sports or music
curl -XGET localhost:9200/megacorp/employee/_search -d '{
  "query": {
    "filtered": {
      "query": {"match_all": {}},
      "filter": {
        "terms": {
          "interests": ["sports", "music"],
          "execution" : "or"
        }
      }
    }
  }
}'

# nested
curl -XGET localhost:9200/megacorp/employee/_search -d '{
  "query": {
    "bool": {
      "must": [
        {
          "nested": {
            "path": "fetish",
            "query": {
              "bool": {
                "must": [ 
                  {"match": {"fetish.name": "RANGE_ROVER_SPORT"}}
                ]
              }
            }
          }
        }
      ]
    }
  }
}'

# has_child - return parent document which have child
curl -XGET localhost:9200/megacorp/employee/_search -d '{
"query": {
    "bool": {
        "should": [
            {
                "has_child" : {
                    "type" : "car",
                    "query" : {
                        "term" : {"name" : "RANGE_ROVER_SPORT"}
                    }
                }
            }
        ]
    }
}
}'

# has parent - return child document which have parent
curl -XGET localhost:9200/megacorp/car/_search -d '{
"query": {
    "bool": {
        "should": [
            {
                "has_parent" : {
                    "type" : "employee",
                    "query" : {
                        "term" : {"last_name" : "Rooney"}
                    }
                }
            }
        ]
    }
}
}'

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

# Search which IF CONDITION
curl -XGET localhost:9200/megacorp/employee/_search?pretty -d '{
    "fields": ["city", "first_name", "last_name", "last_login_at"],
    "post_filter":{
        "script": {
            "script": {
              "script": " if ( _source.city == \"London\" ) { _source.last_login_at == param1 } else { _source.last_login_at == param2 }",
              "params": {"param1" : "2016-01-21", "param2" : "2016-07-11"}
            }
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
    "post_filter": {
        "range": {"age": {"gt" : 35}}
    }
}'

# Simple sort
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "sort" : [{ "city" : "asc" }]
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "fielddata_fields" : ["first_name", "age"]
}'

# order mode
curl -XGET localhost:9200/megacorp/employee/_search?pretty -d '{
    "fields": ["age"],
    "sort" : [{ "age" : {"order" : "asc", "mode" : "avg"} }]
}'
````

````sh
# Custom field
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "script_fields": {"name": {
        "script" : "_source.first_name + _source.last_name"
    }}
}'
````

````sh
curl -XPOST 'localhost:9200/zii/users/18330/_update?pretty' -d '{
"script" : "if (ctx._source.pictures != null) { for (item in ctx._source.pictures) { if (item.picture_id == 3460) { item.type_id = 201201999 } } } "
}'
````

#### Geo queries

````sh
# Find Cristiano Ronaldo by geo_distance
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "query":{
        "bool": {
            "must": {"match_all": {}},
            "filter": {
                "geo_distance" : {
                    "distance" : "100km",
                    "location" : {"lat" : 37, "lon" : 7}
                }
            }
        }
    }
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
````

#### Scroll

````sh
curl 'localhost:9200/zii/users/_search?scroll=1m' -d '{
    "query": {"match_all" : {}}
    "sort" : [{ "last_seen_at" : "desc" }],
    "from": 0,
    "size": 100,
    "_source": true
}'

curl 'localhost:9200/_search/scroll ' -d '{
    "scroll" : "1m", 
    "scroll_id" : "cXVlcnlUaGVuRmV0Y2g7NTsxMzU0MDQzNjY6QnNHMjd0bXlTZ3Ftd1dkblRUd3NQZzsxMDE1NzI4Mjc6ajFzVmtLQUdSaEduRWFRVi1GZE05UTsxMDE1NzI4MjY6ajFzVmtLQUdSaEduRWFRVi1GZE05UTsxMDE1ODAzODc6TURuUG5nbzRUVU9NUUFjSERqM2hIQTsxMDE1ODAzODg6TURuUG5nbzRUVU9NUUFjSERqM2hIQTswOw==" 
}'
````

#### Z examples

````sh
curl -XGET localhost:9200/zii/users/_search -d '{
    "query" : {
        "bool": {
            "filter": {
                "script": {
                    "script": {
                        "inline": "doc['"'"'location'"'"'].arcDistanceInKm(lat, lon)",
                        "params": {"lat": 50.422607185224003, "lon": 30.505957800141999}
                    }
                }
            }
        }
    },
    "post_filter": {
        // "range": {"last_seen_at": {"from": "2016-02-13 14:15:31"}}
        "bool": {
            "must": [
                {
                    "range":{
                        "last_seen_at":{"from":"2016-04-25 16:36:10"}
                    }
                }
            ]
        }
    },
    "from": 0,
    "size": 5
}'

curl -XGET location/zii/users/_search?pretty=true -d '{
    "from" : 10,
    "size" : 10,
    "sort" : [{ "last_seen_at" : "desc" }],
    "query":{
        "bool":{
            "must_not":[
                {"terms": {"user_id": [19194]}}
            ]
        }
    },
    "post_filter": {
        "bool": {
            "must": [
                {"range":{"last_seen_at":{"from": "2016-04-26 14:14:56"} } }
            ]
        }
    },
    "script_fields" : {
        "distance" : {
          "script" : "doc['"'"'location'"'"'].arcDistanceInKm(lat,lon)",
          "params" : {"lat" : 50.422607185224003, "lon" : 30.505957800141999}
        }
    }
}'

curl -XGET localhost:9200/zii/users/_search -d '{
    "from" : 0,
    "size" : 1,
    "fields": ["verification_status_id", "pictures.type_id", "pictures.file_name", "pictures.*"],
    "script_fields": {"e": {"script": "_source.email"}, "s": {"script": "_source.status_id"}}
}'
````
