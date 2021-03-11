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
h=localhost:9200
idx=megacorp
t=employee
url=$h/$idx/$t
url=$h/$idx

# Get document by id 1
curl -XGET $h/$idx/$t/1
curl -XGET $h/$idx/1

# Multi get v2
curl "$h/_mget?pretty" -d '{
    "docs" : [
        {"_index" : "megacorp", "_type" : "employee", "_id" : "1"},
        {"_index" : "megacorp", "_type" : "employee", "_id" : "2"},
        {"_index" : "megacorp", "_type" : "employee", "_id" : "3"}
    ]
}'

# Multi get by certain index and type v2
curl '$h/$idx/$t/_mget?pretty' -d '{
    "docs" : [{"_id" : "1"}, {"_id" : "2"}]
}'
# or v2
curl '$h/$idx/$t/_mget?pretty' -d '{
    "ids" : ["1", "2"]
}'



# find all documents
curl -XGET "$url/_search"
# or
curl "$url/_search?pretty=true&q=*:*"

curl -XGET "$url/_search?q=last_name:Smith"

# validate query
curl -XGET $url_validate/query -d '{
    "query": {"match_all" : {}}
}'

# calculate count of all documents
curl -XGET $url/_count -d '{
    "query": {"match_all" : {}}
}'

# explain
curl -XGET $url/_search -d '{
    "explain": true,
    "query": {"match_all" : {}}
}'
# or
curl -XGET "$url/4/_explain?q=first_name:Louis&pretty"

# profile
curl -XGET $url/_search -d '{
    "profile": true,
    "query": {"match_all" : {}}
}'

# version for each search hit
curl -XGET $url/_search -d '{
    "version": true,
    "query": {"match_all" : {}}
}'

````

