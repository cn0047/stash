Spanner
-

[docs](https://cloud.google.com/spanner/docs/apis)
[query syntax](https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#having_clause)
[golang](https://pkg.go.dev/cloud.google.com/go/spanner)

Spanner - fully managed relational database with unlimited scale and strong consistency.

Spanner instance configured per region.

DB sequences (or auto-increment) - anti-pattern (it creates hotspots), use UUID generator.

Data types:
* bool.
* int64.
* float64.
* numeric.
* string.
* json.
* bytes.
* date.
* timestamp.

* array.

````sql
CREATE TABLE test (
id INT64,
ids ARRAY<INT64>,
names ARRAY<STRING(50)>,
) PRIMARY KEY (id);
````
