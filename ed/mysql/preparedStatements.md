Prepared Statements
-

MySQL 5.0 provides support for server-side prepared statements.

````sql
PREPARE stmt1 FROM 'SELECT ? AS one, ? AS two';
SET @a = 1;
SET @b = 2;
EXECUTE stmt1 USING @a, @b;
+-----+-----+
| one | two |
+-----+-----+
| 1   | 2   |
+-----+-----+
DEALLOCATE PREPARE stmt1;

SET @s = 'SELECT ? AS one, ? AS two';
PREPARE stmt2 FROM @s;
SET @a = 'a';
SET @b = 'b';
EXECUTE stmt2 USING @a, @b;
+-----+-----+
| one | two |
+-----+-----+
| a   | b   |
+-----+-----+
DEALLOCATE PREPARE stmt2;

SET @table = 't1';
SET @s = CONCAT('SELECT * FROM ', @table);
PREPARE stmt3 FROM @s;
EXECUTE stmt3;
DEALLOCATE PREPARE stmt3;

SET @alias = (SELECT Alias FROM Aliases WHERE id = 1);
SET @sql = CONCAT('INSERT INTO ', @alias, ' (somevalue) VALUES (value)');
PREPARE stmt1 FROM @sql;
EXECUTE stmt1;

-- HACKERRANK

-- generate table with incremented numbers
SET @n = 10;
set @i = 1;
SET @t = CONCAT('SELECT 1 AS n', REPEAT(' UNION ALL SELECT @i := @i + 1', @n - 1));
SET @s = CONCAT("SELECT * FROM (", @t, ") t");
PREPARE stmt FROM @s;
EXECUTE stmt;

-- stars
SET @n = 20;
set @i = 1;
SET @t = CONCAT('SELECT 1 AS n', REPEAT(' UNION ALL SELECT @i := @i + 1', @n - 1));
SET @s = CONCAT("SELECT REPEAT('* ', t.n) AS RESULT FROM (", @t, ") t");
PREPARE stmt FROM @s;
EXECUTE stmt;
````

The following SQL statements can be used in prepared statements:
````
ALTER TABLE
CALL
COMMIT
{CREATE | DROP} INDEX
{CREATE | DROP} TABLE
DELETE
DO
INSERT
RENAME TABLE
REPLACE
SELECT
SET
SHOW (most variants)
TRUNCATE TABLE
UPDATE
````
MySQL 5.0.15
````
{CREATE | DROP} VIEW
````
MySQL 5.0.23
````
ANALYZE TABLE
OPTIMIZE TABLE
REPAIR TABLE
````
MySQL 5.1.12
````
CACHE INDEX
CHANGE MASTER
CHECKSUM {TABLE | TABLES}
{CREATE | RENAME | DROP} DATABASE
{CREATE | RENAME | DROP} USER
FLUSH {TABLE | TABLES | TABLES WITH READ LOCK | HOSTS | PRIVILEGES
  | LOGS | STATUS | MASTER | SLAVE | DES_KEY_FILE | USER_RESOURCES}
GRANT
INSTALL PLUGIN
KILL
LOAD INDEX INTO CACHE
RESET {MASTER | SLAVE | QUERY CACHE}
REVOKE
SHOW {AUTHORS | CONTRIBUTORS | WARNINGS | ERRORS}
SHOW BINLOG EVENTS
SHOW CREATE {PROCEDURE | FUNCTION | EVENT | TABLE | VIEW}
SHOW {MASTER | BINARY} LOGS
SHOW {MASTER | SLAVE} STATUS
SLAVE {START | STOP}
UNINSTALL PLUGIN
````
