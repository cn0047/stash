// Cypher administration

Node types: leader, follower.



CALL db.schema();
CALL db.indexes();
CALL db.labels();

CALL dbms.procedures() YIELD name, signature
RETURN *;

CREATE INDEX ON :Person(code);
DROP INDEX ON :Person(code);
CREATE INDEX ON :Person(name, active);

CREATE CONSTRAINT ON (p:Person) ASSERT p.code IS UNIQUE;



// new user
CALL dbms.security.createUser('dbu', 'dbp', false);
