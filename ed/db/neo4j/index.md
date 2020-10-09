Indexes
-

[configuration](https://neo4j.com/docs/operations-manual/current/performance/index-configuration/)

Types: b-tree, full-text.

USING INDEX      - index scan.
USING INDEX SEEK - index seek.
USING JOIN       - @see: https://neo4j.com/docs/cypher-manual/current/query-tuning/using/#_hinting_a_join_on_a_single_node

For `CONTAINS, ENDS WITH` use index provider `lucene+native-3.0` (not `native-btree-1.0`).

`native-btree-1.0` has key size limit 8kB.
`lucene+native-3.0` has key size limit 32kB.

````js
CALL db.indexes();

CREATE INDEX ON :Person(code);
DROP INDEX ON :Person(code);
CREATE INDEX ON :Person(name, active);

// hint
MATCH (p:Person {code: '007'})
USING INDEX p:Person(code)
RETURN p;

MATCH (p:Person {code: '007'})
USING SCAN p:Person // label scan
RETURN p;



// index sampling (statistics and execution plans)
CALL db.resampleIndex(":Person(name)");
CALL db.resampleOutdatedIndexes();
CALL db.clearQueryCaches();
db.prepareForReplanning();


// Examples:
//
// S_IDX - Single-property index.
// S_IDX, STARTS WITH clause - index works.
// S_IDX, ENDS WITH clause   - index works, (search not optimized but it faster than not using index).
// S_IDX, CONTAINS clause    - ↑
//
// C_IDX - Composite index:
// C_IDX, Equality check     - index works if ALL fields from composite index used in query.
// C_IDX, Range comparisons  - ↑
// C_IDX, IN clause          - ↑
// C_IDX, STARTS WITH clause - ↑, like: idx_p1 STARTS WITH 'foo' AND EXISTS (idx_p2)
// C_IDX, ENDS WITH clause   - ↑, like: idx_p1 ENDS WITH 'foo' AND EXISTS (idx_p2)
//                             (search not optimized but it faster than not using index).
// C_IDX, CONTAINS clause    - ↑
// C_IDX, EXISTS clause      - index works for: EXISTS idx_p1 AND EXISTS idx_p2
````

##  Constraint

[see](https://neo4j.com/docs/cypher-manual/current/administration/constraints/#administration-constraints-syntax)

````js
CALL db.constraints();
CREATE CONSTRAINT ON (p:Person) ASSERT p.code IS UNIQUE;
CREATE CONSTRAINT ON (o:Organization) ASSERT o.name IS UNIQUE;

CREATE CONSTRAINT ON (p:Person) ASSERT EXISTS (p.name); // Enterprise Edition
````

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