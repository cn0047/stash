BigQuery
-

[doc](https://cloud.google.com/bigquery/docs)

````sql
select regexp_extract(field, r"^[^\s]+") ua
from tbl`;
````

````sh
curl -XPOST \
  https://bigquery.googleapis.com/bigquery/v2/projects/thisisrealtimelog/datasets/fastly_bigquery/tables/test/insertAll \
  -d '{
    "kind": "bigquery#tableDataInsertAllRequest",
    "rows": [
      {
        "insertId": "1",
        "json": {
          "timestamp": "2019-05-08 16:09:09",
          "time_elapsed": 1,
          "is_tls": "false",
          "client_ip": "194.156.250.183",
          "geo_city": "kyiv",
          "geo_country_code": "ua",
          "request": "post",
          "host": "somehost.com",
          "url": "/blank",
          "request_referer": "/none",
          "request_user_agent": "shell",
          "request_accept_language": "en",
          "request_accept_charset": "utf8",
          "cache_status": "MISS"
        }
      }
    ],
    "skipInvalidRows": false,
    "ignoreUnknownValues": false,
    "templateSuffix": "_kx_"
  }'
````
