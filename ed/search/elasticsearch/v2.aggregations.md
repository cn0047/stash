Aggregations (Analytics)
-

Types:

* methic - stats (avg, cardinality)
* bucketing - categorize into groups
* matrix - may be deleted in future releases
* pipeline - may be deleted in future releases

````sh
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
# same, but ordered by aggregated value (also can use: _term)
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city", "order": {"_count": "asc"}}
    }
  }
}'
# same, but ordered by city
curl -XPOST 'localhost:9200/megacorp/employee/_search?pretty' -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city"}
    }
  },
  "sort" : [{ "city" : {"order" : "desc"} }]
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

# Stats Aggregation
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "stats" : { "stats" : { "field" : "age" } }
    }
}'

# Extended Stats Aggregation
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "e_stats" : { "extended_stats" : { "field" : "age" } }
    }
}'

# Filter Aggregation
curl -XGET localhost:9200/megacorp/employee/_search?pretty -d '{
    "aggs" : {
        "age31" : {
            "filter" : {"range" : {"age": {"gt": 31}}},
            "aggs" : {
                "avg_age" : { "avg" : { "field" : "age" } }
            }
        }
    }
}'

# Missing Aggregation
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "without_age" : { "missing" : { "field" : "age" } }
    }
}'

# Nested Aggregation
# this example won't work it is here just for facilitate future investigations...
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "n_a" : {
            "nested" : { "path" : "location" },
            "aggs" : {
                "m_l" : { "max" : { "field" : "location.lon" } }
            }
        }
    }
}'

# cardinality (distinct)
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "crd" : { "cardinality" : { "field" : "age" } }
    }
}'

# Percentiles
curl -XGET localhost:9200/megacorp/employee/_search -d '{
    "aggs" : {
        "p_stats" : { "percentiles" : { "field" : "age" } }
    }
}'

# get count of each interest
curl -XGET localhost:9200/megacorp/employee/_search?pretty -d '{
  "size": 0,
  "aggs":{
    "interest_count":{"terms":{"field": "interests"}}
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
