Megacorp examples
-

````json
# Create new documents (employee)
curl -XPUT localhost:9200/megacorp/employee/1?routing=JohnSmith -d '{
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
    "interests": [ "forestry", "cars" ]
}'
curl -XPUT localhost:9200/megacorp/employee/4 -d '{
    "first_name" : "Louis",
    "last_name" : "de Funès",
    "age" : 70,
    "about": "Actor. I like movies.",
    "last_login_at": "2012-03-04",
    "city": "Paris",
    "location": {"lat": 48.8567, "lon": 2.3508},
    "interests": [ "fantomas", "theatre", "hollywood" ]
}'
curl -XPUT localhost:9200/megacorp/employee/5 -d '{
    "first_name" : "Cristiano",
    "last_name" : "Ronaldo",
    "age" : 31,
    "about": "Footballer. I like sport.",
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
curl -XPUT localhost:9200/megacorp/employee/13 -d '{
    "first_name" : "Wladimir",
    "last_name" : "Klitschko",
    "age" : 40,
    "about": "REALLY Professional Boxer. Longtime World Heavyweight Champion.",
    "last_login_at": "2016-04-22",
    "city": "Kyiv",
    "location": {"lat": 50.4501, "lon": 30.5234},
    "interests": [ "boxing", "sport", "movie", "hollywood" ]
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

````
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

#### Update

````
# Update particular document
curl -XPOST 'localhost:9200/megacorp/employee/1/_update?pretty' -d '{
  "doc": { "first_name": "JohnnNnn" }
}'

# Update particular document using script
curl -XPOST 'localhost:9200/megacorp/employee/1/_update?pretty' -d '{
  "script" : "ctx._source.age += 100"
}'

# Bulk operations up & del
curl -XPOST 'localhost:9200/megacorp/employee/_bulk?pretty' -d '
{"update": {"_id": "9"}}
{"doc": {"name": "John Doe becomes John DoeeEee"}}
{"delete": {"_id": "8"}}
'

# Scripted update
curl -XPOST 'localhost:9200/megacorp/employee/2/_update' -d '{
    "script" : {
        "inline": "ctx._source.age += count",
        "params" : {"count" : 2}
    }
}'
curl -XPOST 'localhost:9200/megacorp/employee/2/_update' -d '{
    "script" : {
        "inline": "ctx._source.interests += el",
        "params" : {"el" : "rock music"}
    }
}'

# In addition to _source, the following variables are available through the ctx map:
_index, _type, _id, _version, _routing, _parent, _timestamp, _ttl.

# Add new field to certain user
curl -XPOST 'localhost:9200/megacorp/employee/2/_update' -d '{
    "script" : "ctx._source.name_of_new_field = \"value_of_new_field\""
}'

# Update by query - ADD new field
curl -XPOST 'localhost:9200/megacorp/employee/_update_by_query?conflicts=proceed&pretty' -d '{
    "query": {"match_all" : {}},
    "script" : {
        "inline": "ctx._source.likes = \"0\""
    }
}'

# Update by query 1000 documents
curl -XPOST 'localhost:9200/megacorp/employee/_update_by_query&scroll_size=1000' -d '{
  "script": {
    "inline": "ctx._source.likes++"
  }
}'

# Remove a field from the document
curl -XPOST 'localhost:9200/megacorp/employee/2/_update' -d '{
    "script" : "ctx._source.remove(\"name_of_new_field\")"
}'

# Change the operation that is executed.
# This example deletes the doc if the tags field contain blue, otherwise it does nothing (noop):
curl -XPOST 'localhost:9200/megacorp/employee/2/_update' -d '{
    "script" : {
        "inline": "ctx._source.interests.contains(tag) ? ctx.op = \"delete\" : ctx.op = \"none\"",
        "params" : {"tag" : "rock music"}
    }
}'
````

#### Search

````json
# find all employee
curl -XGET localhost:9200/megacorp/employee/_search

curl -XGET localhost:9200/megacorp/employee/_search?q=last_name:Smith

# Calculate count of all documents
curl -XGET localhost:9200/megacorp/employee/_count -d '{
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
    "post_filter": {
        "range": {"age": {"gt" : 35}}
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

# Simple sort
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "sort" : [{ "city" : "asc" }]
}'

curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "fielddata_fields" : ["first_name", "age"]
}'
````
````
# Custom field
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "script_fields": {"name": {
        "script" : "_source.first_name + _source.last_name"
    }}
}'
````

````
curl -XPOST 'localhost:9200/ziipr/users/18330/_update?pretty' -d '{
"script" : "if (ctx._source.pictures != null) { for (item in ctx._source.pictures) { if (item.picture_id == 3460) { item.type_id = 201201999 } } } "
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
# SELECT city, COUNT(*) FROM employee GROUP BY city ORDER BY COUNT(*) DESC
# size=0 to not show search hits
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city"}
    }
  }
}'

# like prev example + AVG(age)
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city"},
      "aggs": {
        "average_age": {
          "avg": {"field": "age"}
        }
      }
    }
  }
}'

# like prev example + ORDER BY average_age DESC
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city", "order": {"average_age": "desc"}},
      "aggs": {
        "average_age": {
          "avg": {"field": "age"}
        }
      }
    }
  }
}'

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
curl -XPUT 'http://localhost:9200/twitter/tweet/1?ttl=1m' -d '{
    "user": "kimchy",
    "message": "Trying out elasticsearch, so far so good?"
}'