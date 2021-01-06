Administration
-

[funcs](https://neo4j.com/docs/operations-manual/current/reference/procedures/)
[conf](https://neo4j.com/docs/operations-manual/current/reference/configuration-settings/)
[import data](https://neo4j.com/docs/operations-manual/current/tools/import/)
[dump](https://neo4j.com/docs/operations-manual/current/tools/dump-load/)
[copy db](https://neo4j.com/docs/operations-manual/current/tools/copy/)
[performance](https://neo4j.com/developer/guide-performance-tuning/)

````sh
systemctl status neo4j

neo4j-admin

dump-config

export NEO4J_dbms_memory_pagecache_size=4G
export NEO4J_dbms_memory_heap_max__size=4G
export EXTENSION_SCRIPT=/extra_conf.sh
````

DBMS - Database Management System.
Instance -  Java process that is running the Neo4j.
Transaction domain - collection of graphs that can be updated within one transaction.
Execution context - query|transaction|internal function|procedure.
Graph - data model within database.

## Cypher

````js
:SYSINFO

SHOW DATABASES;
SHOW DEFAULT DATABASE;
SHOW DATABASE system;
CREATE DATABASE mydb IF not exists; // Enterprise Edition
STOP DATABASE mydb;
START DATABASE mydb;
:USE mydb;
SHOW DATABASE mydb;
DROP DATABASE mydb;
// @see: https://neo4j.com/docs/operations-manual/current/manage-databases/queries/

// thread pool memory usage
CALL dbms.listPools();

CALL dbms.listQueries();
CALL dbms.listActiveLocks(queryId);
CALL dbms.killQuery(queryId);
CALL dbms.killQueries(queryId1, queryId1);

// get running transactions
CALL dbms.listTransactions();

CALL dbms.listConnections();
CALL dbms.killConnection(connectionId);
CALL dbms.killConnections(connectionIds);

CALL dbms.cluster.overview();

CALL db.schema(); // v3
CALL db.schema.visualization();
CALL db.labels();
CALL dbms.procedures() YIELD name, signature RETURN * ;

// properties for node Activity
CALL apoc.meta.schema() YIELD value as schemaMap
UNWIND keys(schemaMap) as label
WITH label, schemaMap[label] as data
WHERE data.type = "node" and label = "Activity"
UNWIND keys(data.properties) as property
WITH label, property, data.properties[property] as propData
RETURN
  label,
  property,
  propData.type as type,
  propData.indexed as isIndexed,
  propData.unique as uniqueConstraint,
  propData.existence as existenceConstraint;

// new user
:USE system;
CALL dbms.security.createUser('dbu', 'dbp', false);
SHOW USERS;

// dynamic settings
CALL dbms.setConfigValue('dbms.logs.query.enabled', '')
// @see: https://neo4j.com/docs/operations-manual/current/configuration/dynamic-settings/

CALL dbms.listConfig()
YIELD name, value
WHERE name STARTS WITH 'dbms.logs'
RETURN name, value;
````

## Export/Import

[import](https://neo4j.com/labs/apoc/4.1/import/)
[export](https://neo4j.com/labs/apoc/4.1/export/)

````js
LOAD CSV WITH HEADERS from 'file:///actors.csv' as row
FIELDTERMINATOR ';'
RETURN row.name;
````

````js
// check import config
CALL dbms.listConfig() YIELD name, value
WHERE name = 'apoc.import.file.use_neo4j_config' RETURN name, value;

````

````js
// export db
CALL apoc.export.csv.all('db.csv', {});
docker exec -it xneo4j sh -c 'ls /var/lib/neo4j/import/db.csv'

// exports csv person
MATCH (p:Person)
WITH collect(p) AS ps
CALL apoc.export.csv.data(ps, [], 'db.persons.csv', {})
YIELD  file, source, format, nodes, relationships, properties, time, rows, batchSize, batches, done, data
RETURN file, source, format, nodes, relationships, properties, time, rows, batchSize, batches, done, data
;



// read from csv
CALL apoc.load.csv('db.csv')
YIELD lineNo, map, list
RETURN *;

// read from csv
CALL apoc.load.csv('db.csv', {skip: 0, limit: 50, header:true, ignore:['status'],
    mapping:{
        name: {type: 'str'},
        vendor: {type: 'str'}
    }
})
YIELD lineNo, map, list
RETURN *;



// import person (⚠️ not working in v4.1.1)
CALL apoc.import.csv([{fileName: 'db.persons.csv', labels: ['Person']}], [], {});
CALL apoc.import.csv([{fileName: 'file://db.persons.csv', labels: ['Person']}], [], {});
CALL apoc.import.csv(
  [{fileName: 'db.persons.csv', labels: ['Person']}],
  [],
  {delimiter: ',', arrayDelimiter: ';', stringIds: false}
);

// import (✅ works)
CALL apoc.periodic.iterate(
    'CALL apoc.load.csv("db.persons.csv") yield map as row return row ',
    'CREATE (p:Person) SET p = row',
    {batchSize: 50, iterateList: true, parallel: true}
);
````

````js
// export db
CALL apoc.export.json.all('db.json', {});
docker exec -it xneo4j sh -c 'ls /var/lib/neo4j/import/db.json'

// exports json person
MATCH (p:Person)
WITH collect(p) AS ps
CALL apoc.export.json.data(ps, [], 'db.persons.json', {})
YIELD  file, source, format, nodes, relationships, properties, time, rows, batchSize, batches, done, data
RETURN file, source, format, nodes, relationships, properties, time, rows, batchSize, batches, done, data
;

// read json db
CALL apoc.load.json('db.json')
YIELD value
RETURN value;

// read json persons
CALL apoc.load.json('db.persons.json')
YIELD value
RETURN value.properties.name;

// import (✅ works)
CALL apoc.load.json('db.persons.json')
YIELD value
MERGE (p:Person {name: value.properties.name})
SET p.imported = true;

````

Import separate CSV files:

````sh
# 1) ⚠️ delete db dir, import won't work if dir is not empty
rm -rf .data/.docker/neo4j_41

# 2) run docker and don't start db (with /bin/bash)

# 3)
docker exec -it xneo4j sh -c '
  echo "pID:ID,name,:LABEL" > import/person.h.csv ;
  echo "blofeld,Blofeld,:Person" > import/person.csv ;
  echo "bond,Bond,:Person" >> import/person.csv ;

  echo "cID:ID,name,vendor,:LABEL" > import/car.h.csv ;
  echo "silverwraith,Silver Wraith,Rolls-Royce,:Car" > import/car.csv ;
  echo "db11v8,DB11 V8,Aston Martin,:Car" >> import/car.csv ;

  echo ":START_ID,val,:END_ID,:TYPE" > import/rel.h.csv ;
  echo "blofeld,1,silverwraith,:LIKES" > import/rel.csv ;
  echo "bond,2,db11v8,:LIKES" >> import/rel.csv ;

  neo4j-admin import \
    --database mydb \
    --nodes import/person.h.csv,import/person.csv \
    --nodes import/car.h.csv,import/car.csv \
    --relationships import/rel.h.csv,import/rel.csv \
'
````

Dump:

````sh
# ) run docker and don't start db (with /bin/bash)

# )
docker exec -it xneo4j sh -c '
  neo4j-admin dump --database=mydb --to=1.dump
'

docker exec -it xneo4j sh -c '
  neo4j-admin load --from=1.dump --database=mydb2
'
````

## Configuration

````sh
$NEO4J_HOME/conf/neo4j.conf
$NEO4J_HOME/data             # data dir
$NEO4J_HOME/import           # `LOAD CSV`
$NEO4J_HOME/plugins
$NEO4J_HOME/logs

echo $NEO4J_HOME
echo $NEO4J_CONF

dbms.directories.data=data
dbms.directories.plugins=plugins
dbms.directories.logs=logs
dbms.directories.lib=lib
dbms.directories.run=run
dbms.directories.metrics=metrics

dbms.memory.heap.initial_size=
dbms.memory.heap.max_size=
dbms.jvm.additional=

dbms.default_advertised_address=localhost
dbms.default_database=neo4j
dbms.default_listen_address=localhost
# @see: https://neo4j.com/docs/operations-manual/current/configuration/connectors/

# ports
# @see: https://neo4j.com/docs/operations-manual/current/configuration/ports/

````
