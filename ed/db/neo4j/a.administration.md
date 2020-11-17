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
CALL dbms.killConnections(connectionIds);
CALL dbms.killConnection(connectionId);

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

## Configuration

````sh
<neo4j-home>/conf/neo4j.conf
<neo4j-home>/data             # data dir
<neo4j-home>/import           # `LOAD CSV`
<neo4j-home>/plugins
<neo4j-home>/logs

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
