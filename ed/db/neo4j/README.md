neo4j
-

[docs](https://neo4j.com/docs/)

Cypher - declarative graph query language.
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
````
