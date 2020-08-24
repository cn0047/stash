neo4j
-
<br>4.1
<br>3.5

[docs](https://neo4j.com/docs/)
[data types](https://neo4j.com/docs/cypher-manual/current/syntax/values/)
[funcs](https://neo4j.com/docs/cypher-manual/current/functions/)
[keywords](https://neo4j.com/docs/cypher-manual/current/keyword-glossary/)
[cypher styleguide](https://neo4j.com/docs/cypher-manual/current/styleguide/)

````sh
cypher-shell -u neo4j -p test
cypher-shell -a neo4j:test@localhost:7687
````

Bloom - graph exploration app.

Node - represents entity.
Label - describes the domain `:user :suspended`.
Relationship - connects two nodes, it always have a direction.
Properties - key/value pairs used to add qualities to nodes and relationships.
Traversing - means visiting nodes by following relationships according to some rules.
Schema - indexes and constraints.

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
UNWIND - expands list into sequence of rows.

## Performance

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
