Foreign key
-

````
[CONSTRAINT [symbol]] FOREIGN KEY
    [index_name] (index_col_name, ...)
    REFERENCES tbl_name (index_col_name,...)
    [ON DELETE reference_option]
    [ON UPDATE reference_option]

reference_option:
    RESTRICT | CASCADE | SET NULL | NO ACTION
````

The FOREIGN KEY clause is specified in the child table.
The parent and child tables must use the same storage engine.
Corresponding columns in the foreign key and the referenced key must have similar data types
(the size and sign of integer types must be the same, the length of string types need not be the same,
BLOB and TEXT columns cannot be included in a foreign key).

MySQL requires indexes on foreign keys and referenced keys
so that foreign key checks can be fast and not require a table scan.

InnoDB does not currently support foreign keys for tables with user-defined partitioning.
This includes both parent and child tables.

#### Referential Actions

MySQL rejects any INSERT or UPDATE operation that attempts to create a foreign key value in a child table
if there is no a matching candidate key value in the parent table.

When an UPDATE or DELETE operation affects a key value in the parent table
that has matching rows in the child table,
the result depends on the referential action specified using ON UPDATE and ON DELETE
subclauses of the FOREIGN KEY clause.

*CASCADE:*
Delete or update the row from the parent table,
and automatically delete or update the matching rows in the child table.
Both ON DELETE CASCADE and ON UPDATE CASCADE are supported.
(Currently, cascaded foreign key actions do not activate triggers.)

*SET NULL:*
Delete or update the row from the parent table,
and set the foreign key column or columns in the child table to NULL.
Both ON DELETE SET NULL and ON UPDATE SET NULL clauses are supported.
(If you specify a SET NULL action,
make sure that you have not declared the columns in the child table as NOT NULL.)

*RESTRICT:*
Rejects the delete or update operation for the parent table.
Specifying RESTRICT (or NO ACTION) is the same as omitting the ON DELETE or ON UPDATE clause.

*NO ACTION:*
Equivalent to RESTRICT.

*SET DEFAULT:*
InnoDB rejects table definitions containing ON DELETE SET DEFAULT or ON UPDATE SET DEFAULT clauses.

````sql
-- Shows foreign keys
SELECT
    i.TABLE_SCHEMA,
    i.TABLE_NAME,
    i.CONSTRAINT_TYPE,
    i.CONSTRAINT_NAME,
    k.REFERENCED_TABLE_NAME,
    k.REFERENCED_COLUMN_NAME
FROM information_schema.TABLE_CONSTRAINTS i
LEFT JOIN information_schema.KEY_COLUMN_USAGE k ON i.CONSTRAINT_NAME = k.CONSTRAINT_NAME
WHERE i.CONSTRAINT_TYPE = 'FOREIGN KEY' AND i.TABLE_SCHEMA = 'dbName'
;

CREATE TABLE parent (
    id INT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;
CREATE TABLE child (
    id INT KEY,
    parent_id INT,
    INDEX par_ind (parent_id),
    FOREIGN KEY (parent_id) REFERENCES parent(id) ON DELETE CASCADE
) ENGINE=INNODB;

SELECT
    i.TABLE_SCHEMA,
    i.TABLE_NAME,
    i.CONSTRAINT_TYPE,
    i.CONSTRAINT_NAME,
    k.REFERENCED_TABLE_NAME,
    k.REFERENCED_COLUMN_NAME
FROM information_schema.TABLE_CONSTRAINTS i
LEFT JOIN information_schema.KEY_COLUMN_USAGE k ON i.CONSTRAINT_NAME = k.CONSTRAINT_NAME
WHERE i.CONSTRAINT_TYPE = 'FOREIGN KEY' AND i.TABLE_SCHEMA = 'test'
;
+--------------+------------+-----------------+-----------------+-----------------------+------------------------+
| TABLE_SCHEMA | TABLE_NAME | CONSTRAINT_TYPE | CONSTRAINT_NAME | REFERENCED_TABLE_NAME | REFERENCED_COLUMN_NAME |
+--------------+------------+-----------------+-----------------+-----------------------+------------------------+
| test         | child      | FOREIGN KEY     | child_ibfk_1    | parent                | id                     |
+--------------+------------+-----------------+-----------------+-----------------------+------------------------+

INSERT INTO parent VALUES (1), (2), (3);
INSERT INTO child VALUES (1, 1), (2, 2), (3, 3);

SELECT * FROM parent;
+----+
| id |
+----+
|  1 |
|  2 |
|  3 |
+----+
SELECT * FROM child;
+------+-----------+
| id   | parent_id |
+------+-----------+
|    1 |         1 |
|    2 |         2 |
|    3 |         3 |
+------+-----------+

DELETE FROM parent WHERE id = 2;
SELECT * FROM parent;
+----+
| id |
+----+
|  1 |
|  3 |
+----+
SELECT * FROM child;
+------+-----------+
| id   | parent_id |
+------+-----------+
|    1 |         1 |
|    3 |         3 |
+------+-----------+

SELECT *
FROM parent
JOIN child ON parent.id = child.parent_id
;

DELETE FROM child
WHERE parent_id IN (SELECT id FROM parent)
;

DROP TABLE child;
DROP TABLE parent;
````
