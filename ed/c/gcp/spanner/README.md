Spanner
-

[docs](https://cloud.google.com/spanner/docs/apis)
[query syntax](https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#having_clause)
[funcs](https://cloud.google.com/spanner/docs/reference/standard-sql/syntax)
[mutations](https://cloud.google.com/spanner/docs/modify-mutation-api)
[golang](https://cloud.google.com/spanner/docs/getting-started/go)
[golang](https://pkg.go.dev/cloud.google.com/go/spanner)

DDL - data definition language.
DML - data manipulation language.

````sh
dbi=test-instance
db=test-db
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



# emulator
gcloud emulators spanner start

# start docker
# image: gcr.io/cloud-spanner-emulator/emulator:latest
# stop
docker stop `docker ps | grep spanner-emulator | awk '{print $1}'`

export SPANNER_EMULATOR_HOST=localhost:9010
#
gcloud config set auth/disable_credentials true
gcloud config set api_endpoint_overrides/spanner "http://localhost:9020/"



cfg=test
prj=test-project
dbi=test-instance
db=test-db

gcloud spanner instances list
# emulator
gcloud spanner instances create $dbi \
  --config=$cfg --description="TestEmulatorInstance" --nodes=1

gcloud spanner databases list --instance=$dbi
# create db
gcloud spanner databases create $db --instance=$dbi

ddl 'CREATE TABLE test (id INT64 NOT NULL, msg STRING(100), data JSON) PRIMARY KEY(id)'
ddl 'DROP TABLE test'
q "INSERT INTO test (id, msg) VALUES (1, 'one')"
q 'SELECT * FROM test'

````

Spanner instance can be per region or multi-regional.
Spanner - fully managed relational database with unlimited scale and strong consistency.
Spanner DB dialect: Google Standard SQL, PostgreSQL.
DB sequences (or auto-increment) - anti-pattern (it creates hotspots), use UUID generator.

[Data types](https://cloud.google.com/spanner/docs/reference/standard-sql/data-types):
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
id STRING(36) NOT NULL,
i1 INT64 NOT NULL,
s1 STRING(100) NOT NULL,
b1 BOOL NOT NULL,
d1 DATE NOT NULL,
ts TIMESTAMP NOT NULL,

ids ARRAY<INT64>,
names ARRAY<STRING(50)>,
) PRIMARY KEY (id1, id2);

SELECT 2 IN UNNEST(ARRAY_CONCAT([1, 2], [3, 4])) in_array;

````
