MySql
-

*5.1.72-rel14.10*

####Options
````sql
SET SQL_BIG_SELECTS = 1;
SET SQL_SAFE_UPDATES = 0;
````

````sql
ALTER TABLE table ADD UNIQUE KEY (field);
ALTER TABLE table DROP COLUMN field;
DROP INDEX keyName ON table;
````

####Flashback
````sql
-- First works AND than OR

LIKE '[JM]%' -- begins on J or M
% - *
_ - 1 char
[] - group of symblols

COLLATE UTF8_GENERAL_CI LIKE

SELECT COALESCE(field, 0) FROM table; -- if field is null returns 0

OREDER BY ASC -- default ASC

LEFT OUTER JOIN -- search rows that not exists in second table

INSERT INTO table VALUES (default); -- default value

UPDATE tab SET field = DEFAULT(field);

CHECK TABLE tab;
ANALYZE TABLE tab;
REPAIR TABLE tab;
OPTIMIZE TABLE tab;

SELECT VERSION();
SELECT USER();
SHOW GRANTS FOR 'usr';
SHOW GRANTS;
SHOW PRIVILEGES;

SHOW FULL TABLES;

Масштабирование по вертикали - повышение производительности сервера.
Масштабирование по горезонтали - распределение данных по нескольким серверам (репликация).
````

````sql
SELECT SQL_CACHE
SELECT SQL_NO_CACHE
SELECT SQL_CALC_FOUND_ROWS * FROM table;
SELECT FOUND_ROWS();

SHOW COLUMNS FROM table LIKE '%';

INSERT INTO tbl1 (field1) SELECT field2 FROM tbl2;

SELECT AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'databaseName' AND TABLE_NAME = 'tableName';

SELECT SUM(field) FROM table GROUP BY field WITH ROLLUP;
````

####Functions
````sql
ROW_COUNT() -- after insert, update, delete
LAST_INSERT_ID()
CONNECTION_ID()

UNIX_TIMESTAMP()
CURTIME()
sysdate() -- now
LAST_DAY(date) -- last day in month

REPLACE('www.site.com', 'w', 'w')

SELECT ELT(1, 'foo', 'bar');       -- foo
SELECT FIELD('foo', 'foo', 'bar');  -- 1
SELECT FIND_IN_SET('b', 'a,b,cd'); -- 2

INET_ATON(ip)
INET_NTOA(i)

UUID()
````

####Tricks
````
~/.mysql_history

tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
````
