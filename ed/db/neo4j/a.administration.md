Administration
-

````sh
systemctl status neo4j

dump-config

export NEO4J_dbms_memory_pagecache_size=4G
export NEO4J_dbms_memory_heap_max__size=4G
export EXTENSION_SCRIPT=/extra_conf.sh
````

## Cypher

Node types: leader, follower.

````js
SHOW DATABASES;
SHOW DATABASE system;
CREATE DATABASE mydb if not exists; // Enterprise Edition
STOP DATABASE mydb;
START DATABASE mydb;
SHOW DATABASE mydb;
DROP DATABASE mydb;

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
````
