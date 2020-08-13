// Indexes

CALL db.indexes();

CREATE INDEX ON :Person(code);
DROP INDEX ON :Person(code);
CREATE INDEX ON :Person(name, active);

// constraint
// @see: https://neo4j.com/docs/cypher-manual/current/administration/constraints/#administration-constraints-syntax
CALL db.constraints();
CREATE CONSTRAINT ON (p:Person) ASSERT p.code IS UNIQUE;
// CREATE CONSTRAINT ON (p:Person) ASSERT EXISTS (p.name); // Enterprise Edition

// fulltext index
// @see: https://neo4j.com/docs/cypher-manual/current/administration/indexes-for-full-text-search
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
