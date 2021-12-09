Spanner
-

[docs](https://cloud.google.com/spanner/docs/apis)
[query syntax](https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#having_clause)
[golang](https://cloud.google.com/spanner/docs/getting-started/go)
[golang](https://pkg.go.dev/cloud.google.com/go/spanner)

````sh
gcloud emulators spanner start

# emulator
export SPANNER_EMULATOR_HOST=localhost:9010
#
gcloud config set auth/disable_credentials true
gcloud config set api_endpoint_overrides/spanner http://localhost:9020/



gcloud spanner instances list
# emulator
gcloud spanner instances create test-instance \
  --config=to --description="TestEmulatorInstance" --nodes=1

gcloud spanner databases list
# create db
gcloud spanner databases create testdb --instance=test-instance

gcloud spanner databases ddl update testdb --instance=test-instance \
  --ddl='CREATE TABLE test (id INT64 NOT NULL, msg STRING(100) ) PRIMARY KEY(id);'

````

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

SELECT 2 IN UNNEST(ARRAY_CONCAT([1, 2], [3, 4])) in_array;

````
