# Nested objects limit.



h=localhost:9200
idx=nested_test

# put mapping
curl -XPUT $h/$idx -H $jh -d '{
  "settings": {
    "index.mapping.nested_objects.limit": 10,
    "number_of_shards": 8
  },
  "mappings": {
    "properties": {
      "id": {
        "type": "keyword"
      },
      "name": {
        "type": "text"
      },
      "jobs": {
        "type": "nested",
        "properties": {
          "id": {
            "type": "keyword"
          },
          "name": {
            "type": "text"
          }
        }
      }
    }
  }
}'

# put doc
curl -XPOST $h/$idx/_doc/1 -H $jh -d '{
  "id": "1",
  "jobs": [
    {"id": "j1"}
  ]
}'

# search
curl "$h/$idx/_search" | jq > /tmp/x.json
curl "$h/$idx/_search" | jq '.hits.hits[0]._source.jobs' > /tmp/x.json
curl "$h/$idx/_search" | jq '.hits.hits[0]._source.jobs | length'

# put nested element into array in loop
i=1
while true; do
  i=$(($i+1))
  echo $i
  curl -XPOST "$h/$idx/_update/1" -H $jh -d '{
      "script" : {
          "inline": "ctx._source.jobs.add(params.job)",
          "params" : {"job" : {"id": "j'$i'"}}
      }
  }' | jq
done

# delete with threshold
curl -s -XPOST "$h/$idx/_update/1" -H $jh -d '{
    "script" : {
        "inline": "if (ctx._source.jobs.length >= params.threshold) { ctx._source.jobs.remove(0); }",
        "params" : {"job" : {"id": "j11"}, "threshold": 10}
    }
}' | jq

