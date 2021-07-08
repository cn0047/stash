Neo4j
-
<br>4.1
<br>3.5

[docs](https://neo4j.com/docs/)
[data types](https://neo4j.com/docs/cypher-manual/current/syntax/values/)
[funcs](https://neo4j.com/docs/cypher-manual/current/functions/)
[keywords](https://neo4j.com/docs/cypher-manual/current/keyword-glossary/)
[cypher styleguide](https://neo4j.com/docs/cypher-manual/current/styleguide/)
[algorithms](https://neo4j.com/developer/graph-data-science/graph-algorithms/)
[links](https://neo4j.com/developer/resources/)

````sh
cypher-shell -u neo4j -p test
cypher-shell -a neo4j:test@localhost:7687
````

Neo4j - is ACID graph DB.
Neo4j sandbox - temporary neo4j instance in cloud, expire in 3 days.
Bloom - graph exploration app.
ETL - extracts schema from relational db and turn it into graph schema.
Bolt - efficient, lightweight client-server protocol for db apps, (default port 7687).

Node - represents entity.
Dense node - node with huge number of relationships.
Label - describes the domain `:user :suspended`.
Relationship - connects two nodes, it always have a direction.
Properties - key/value pairs used to add qualities to nodes and relationships.
Traversing - means visiting nodes by following relationships according to some rules.
Schema - indexes and constraints.

Connections:
`neo4j://` - will have server-side routing enabled,
`bolt://` - will not.

````sh
curl http://localhost:7474/db/data
curl http://localhost:7474/db/data/batch
curl http://localhost:7474/db/data/node/1

system # the system database
neo4j  # the default database
````

## Transaction

* Auto-commit transaction - is invoked using the `session.Run` method.
* Transaction function - transactional units of work (`session.WriteTransaction`).
* Explicit transactions - `BEGIN, COMMIT and ROLLBACK` operations.

Transactions can be executed in: read or write mode.
Access modes can be supplied: per transaction or per session.

## Cypher

Cypher - declarative graph query language.

````
=~ 'regexp'
(?i) case-insensitive
(?m) multiline
````

WITH -  makes query parts chained, piping the results 1 part to 2nd.
OPTIONAL MATCH - return nulls for missing parts of the pattern.

## Performance

EXPLAIN ... - see the execution plan but not run.
PROFILE ... - run and see which operators are doing most of the work.

`WHERE p.name STARTS WITH ''` - may improve performance.
`WITH p` - berore RETURN, may improve performance.

````js
// replan=default
// replan=force   - force a replan.
// replan=skip    - if valid plan already exists, it will be used.
CYPHER replan=force MATCH ...
CYPHER replan=force EXPLAIN MATCH ...
````

[In query plan](https://neo4j.com/docs/cypher-manual/current/execution-plans/operator-summary/):
* AllNodesScan
* NodeByLabelScan
* NodeIndexScan
* NodeIndexSeek - Single-property index, Equality check.
* NodeIndexSeekByRange
* NodeIndexContainsScan
* NodeIndexEndsWithScan

Execution Plans:
* Evaluation model.
* Eager and lazy evaluation.

Rows - rows number that operator produced.
EstimatedRows - estimated rows number that is expected to be produced.
DbHits - abstract unit of storage engine work.
Page Cache Hits, Page Cache Misses, Page Cache Hit Ratio - page cache
used to cache data and avoid accessing disk,
so having high number of hits and low number of misses make query faster.

## Ops

Hardware requirements (cloud):
CPU - 2vCPU min.
Memory - 2GB min.
Storage - 10GB min.
JVM.

## Causal Cluster

Node types: Leader, Follower.

Discovery protocol takes information from `causal_clustering.initial_discovery_members` conf.
Core Servers (can have multiple roles, one for each database).
Raft protocol.
Read Replica (not involved in the Raft protocol).
Whiteboard (routing & monitoring) - provides a view of all live Read Replicas.
Raft Leader - maintain own logs with respect to the current Leader's log.
Raft Follower.
Raft Candidate.
Catchup protocol.
Read Replica shutdown - read replica invoke discovery protocol to remove itself from whiteboard.
Core shutdown - handled via Raft protocol.

````sh
causal_clustering.minimum_core_cluster_size_at_formation=F
# min number of core machines initially required to start a cluster

causal_clustering.minimum_core_cluster_size_at_runtime=R

````
Good practice when F = R (↑).

Quorum (majority) - NumCoresRunning/2 + 1.
