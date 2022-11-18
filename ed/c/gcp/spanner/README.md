Spanner
-

[docs](https://cloud.google.com/spanner/docs/apis)
[query syntax](https://cloud.google.com/spanner/docs/reference/standard-sql/query-syntax#having_clause)
[funcs](https://cloud.google.com/spanner/docs/reference/standard-sql/syntax)
[mutations](https://cloud.google.com/spanner/docs/modify-mutation-api)
[TTL](https://cloud.google.com/spanner/docs/ttl/working-with-ttl#syntax)
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
# docker image: gcr.io/cloud-spanner-emulator/emulator:latest
cfg=test
prj=test-project
dbi=test-instance
db=test-db
# start
gcloud config configurations activate $cfg || gcloud config configurations create $cfg
gcloud config set project $prj
gcloud config set auth/disable_credentials true
gcloud config set api_endpoint_overrides/spanner 'http://localhost:9020/'
gcloud -q emulators spanner start &
gcloud -q spanner instances create $dbi --config=$cfg --description="TestEmulatorInstance" --nodes=1
gcloud -q spanner databases create $db --instance=$dbi
# stop
docker stop `docker ps | grep spanner-emulator | awk '{print $1}'`

# !!!
export SPANNER_EMULATOR_HOST=localhost:9010



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
CREATE TABLE fields_test (
id1 STRING(36) NOT NULL,
id2 STRING(36) NOT NULL,
b1  BOOL,
i1  INT64,
f1  FLOAT64,
n1  NUMERIC,
s1  STRING(100),
j1  JSON,
bt1 BYTES(MAX),
d1  DATE,
t1  TIMESTAMP,
t2  TIMESTAMP OPTIONS (allow_commit_timestamp=true),
a1  ARRAY<INT64>,
a2  ARRAY<STRING(50)>
) PRIMARY KEY (id1, id2);

SELECT 2 IN UNNEST(ARRAY_CONCAT([1, 2], [3, 4])) in_array;

````
