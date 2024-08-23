View
-

View - virtual table that does not store data.
When query view - DB executes the underlying query each time to retrieve data.
Do not provide any performance benefits,
performance depends on efficiency of underlying query.

Materialized View - physical copy of query result at the time it was created or refreshed.
Data in MV not automatically updated when the underlying data changes.
Can significantly improve query performance
for complex or resource-intensive queries, because MV store precomputed results.

````sql
CREATE TABLE test (
  id SERIAL NOT NULL PRIMARY KEY,
  data TEXT
);
INSERT INTO test VALUES (1, 'one'), (2, 'two');
SELECT * FROM test;

CREATE VIEW vtest AS SELECT * FROM test;
CREATE MATERIALIZED VIEW mvtest AS SELECT * FROM test;
SELECT * FROM vtest;
SELECT * FROM mvtest;
 id | data
----+------
  1 | one
  2 | two

INSERT INTO test VALUES (3, 'three');

SELECT * FROM vtest;
 id | data
----+-------
  1 | one
  2 | two
  3 | three

SELECT * FROM mvtest;
 id | data
----+------
  1 | one
  2 | two

REFRESH MATERIALIZED VIEW mvtest; -- full data delete and full data populate
````
