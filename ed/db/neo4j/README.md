neo4j
-

[docs](https://neo4j.com/docs/)
[data types](https://neo4j.com/docs/cypher-manual/current/syntax/values/)

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
