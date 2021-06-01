Full Text Search
-

Neo4j is using lucene engine.

## Fulltext index

[see](https://neo4j.com/docs/cypher-manual/current/administration/indexes-for-full-text-search)

````js
// createNodeIndex | createRelationshipIndex
CALL db.index.fulltext.createNodeIndex("idx_name", ["Person", "Organization"], ["name"])
CALL db.index.fulltext.createNodeIndex("idx_code", ["Person"], ["code"], {analyzer: "url_or_email", eventually_consistent: "true"})
CALL db.index.fulltext.listAvailableAnalyzers;
CALL db.index.fulltext.drop("idx_code")

// fulltext search
CALL db.index.fulltext.queryNodes("name", "bond") YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes("name", "*ond") YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes("name", "bon*") YIELD node, score RETURN node, score;
CALL db.index.fulltext.queryNodes("name", "b*nd") YIELD node, score RETURN node, score;
````
