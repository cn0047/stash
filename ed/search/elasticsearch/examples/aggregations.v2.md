Aggregations (Analytics)
-

Types:
* methic - stats (avg, cardinality)
* bucketing - categorize into groups
* matrix - may be deleted in future releases
* pipeline - may be deleted in future releases

````sh
url=localhost:9200/megacorp/employee

# SELECT city, COUNT(*) FROM employee GROUP BY city ORDER BY COUNT(*) DESC
# size=0 to not show search hits
curl -XPOST "$url/_search?pretty" -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city"}
    }
  }
}'
# Same, but ordered by aggregated value (also can use: _term)
curl -XPOST "$url/_search?pretty" -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city", "order": {"_count": "asc"}}
    }
  }
}'
# Same, but ordered by city
curl -XPOST "$url/_search?pretty" -d '{
  "size": 0,
  "aggs": {
    "group_by_city": {
      "terms": {"field": "city"}
    }
  },
  "sort" : [{ "city" : {"order" : "desc"} }]
}'

# Like prev example + AVG(age)
curl -XPOST "$url/_search?pretty" -d '{
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

# Like prev example + ORDER BY average_age DESC
curl -XPOST "$url/_search?pretty" -d '{
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
curl -XGET $url/_search -d '{
    "aggs" : {
        "stats" : { "stats" : { "field" : "age" } }
    }
}'

# Extended Stats Aggregation
curl -XGET $url/_search -d '{
    "aggs" : {
        "e_stats" : { "extended_stats" : { "field" : "age" } }
    }
}'

# Filter Aggregation
curl -XGET "$url/_search?pretty" -d '{
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
curl -XGET $url/_search -d '{
    "aggs" : {
        "without_age" : { "missing" : { "field" : "age" } }
    }
}'

# Nested Aggregation
# this example won't work it is here just for facilitate future investigations...
curl -XGET $url/_search -d '{
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
curl -XGET $url/_search -d '{
    "aggs" : {
        "crd" : { "cardinality" : { "field" : "age" } }
    }
}'

# Percentiles
curl -XGET $url/_search -d '{
    "aggs" : {
        "p_stats" : { "percentiles" : { "field" : "age" } }
    }
}'

# get count of each interest
curl -XGET "$url/_search?pretty" -d '{
  "size": 0,
  "aggs":{
    "interest_count":{"terms":{"field": "interests"}}
  }
}'

curl -XGET $url/_search -d '{
    "query": {
        "match": { "last_name": "smith" }
    },
    "aggs": {
        "all_interests": {
            "terms": { "field": "interests" }
        }
    }
}'

curl -XGET $url/_search -d '{
    "aggs" : {
        "all_interests" : {
            "terms" : { "field" : "interests" },
            "aggs" : {
                "avg_age" : { "avg" : { "field" : "age" } }
            }
        }
    }
}'



url=$h/$idx
curl -XPOST "$url/_search?size=0" -H $ctj -d '{
    "aggs" : {
        "aggregated" : {
            "filter" : {
              "bool": {"must": [
                {"match": {"name": "'$m'"}},
                {"term": {"properties.by_user_id": "'$uId'"}},
                {"range" : {"timestamp": {"from": 1178323112, "to": 2579519655}}}
              ]}
            },
            "aggs": {
                "aggregated": {"avg": {"field" : "properties.value"}}
            }
        }
    }
}' | jq
````
