// Cypher administration



CALL db.schema();
CALL db.indexes();

CREATE INDEX ON :Person(code);
DROP INDEX ON :Person(code);
CREATE INDEX ON :Person(name, active);

CREATE CONSTRAINT ON (p:Person) ASSERT p.code IS UNIQUE;
