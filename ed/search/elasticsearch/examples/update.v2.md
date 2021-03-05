Update
-

The update operation supports the following query-string
[parameters](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-update.html#_parameters_3).

````sh
# Update particular document and particular property
curl -XPOST 'localhost:9200/megacorp/employee/1/_update?pretty' -d '{
  "doc": { "first_name": "JohnnNnn" }
}'
# another example
curl -XPOST 'localhost:9200/megacorp/employee/2/_update?pretty' -d '{
  "doc": { "first_name": "JJJane" }
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



# just example
def found = -1;
  for (int i = 0; i < ctx._source.arr.length; ++i) {
    if (ctx._source.arr[i].id==params.id) {
      found=i;
    }
  }
  if (ctx._source.id == "" || ctx._source.id.compareTo(params.id) > 0 ) {
    ctx._source.id = params.id;
  } else if (ctx._source.id == params.id && ctx._source.rev <= params.rev ) {
    ctx._source.id = params.id;
  }
  if (found==-1) {
    ctx._source.arr.add(params.id);
  } else if (ctx._source.arr[found].rev < params.rev) {
    ctx._source.arr[found]=params
  }
  long seven_days_ago = Instant.ofEpochMilli(params.now).minus(7, ChronoUnit.DAYS).toEpochMilli();
  ctx._source.arr.removeIf(doc -> doc.created_at_ms != null && doc.created_at_ms < seven_days_ago);
````
