Spanner
-

[docs](https://cloud.google.com/spanner/docs/apis)
[query syntax](https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#having_clause)
[funcs](https://cloud.google.com/spanner/docs/reference/standard-sql/syntax)
[golang](https://cloud.google.com/spanner/docs/getting-started/go)
[golang](https://pkg.go.dev/cloud.google.com/go/spanner)

````sh
db=test-db
dbi=test-instance
ddl() {
  gcloud spanner databases ddl update $db --instance=$dbi --ddl=$1
}
ddl2() {
  s=`echo $1 | tr '\n' ' '`
  ddl $s
}
q() {
  gcloud spanner databases execute-sql $db --instance=$dbi --sql=$1
}



gcloud emulators spanner start

# emulator
export SPANNER_EMULATOR_HOST=localhost:9010
#
gcloud config set auth/disable_credentials true
gcloud config set api_endpoint_overrides/spanner http://localhost:9020/



db=test-db
dbi=test-instance

gcloud spanner instances list
# emulator
gcloud spanner instances create $dbi \
  --config=to --description="TestEmulatorInstance" --nodes=1

gcloud spanner databases list --instance=$dbi
# create db
gcloud spanner databases create $db --instance=$dbi

ddl 'CREATE TABLE test (id INT64 NOT NULL, msg STRING(100), data JSON) PRIMARY KEY(id)'
ddl 'DROP TABLE test'
q "INSERT INTO test (id, msg) VALUES (1, 'one')"
q 'SELECT * FROM test'

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
