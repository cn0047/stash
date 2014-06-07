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
ALTER TABLE table MODIFY field VARCHAR(7) NOT NULL DEFAULT '';
DROP INDEX keyName ON table;

SHOW TABLE STATUS WHERE name = 'table'; # Info about table, with creation date.

SELECT
    table_name, create_time
FROM information_schema.tables
WHERE table_schema='aff_easydate_biz' AND create_time BETWEEN '2014-05-05' AND '2014-05-07';
````

####mySqlDump
````
mysqldump -h hostname -u user -d --skip-triggers --single-transaction --complete-insert --extended-insert --quote-names dbname table |gzip > sql.gz
mysqldump -hHost Base table | gzip | uuencode table.gz | mail mail@com.com -s table
mysqldump -h Host Base table --where="id=11" | mail mail@com.com
php-mysqldump -h HOST -u USER -a table -f sql -e "SELECT * FROM table LIMIT 10"
mysql -h Host -D Base -te "SELECT NOW()" | mail mail@com.com
mysql -hHost -uUser -pPass base ~ table
mysql -h Host -D Base -e "select * from table where id in (1,2);" | gzip | pv | uuencode result.csv.gz | mail mail@com.com -s "query result"
zcat table.gz | host -uUser -pPass base
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
Statement based replication - на мастере выполнился запрос - на слейвы приходит такой же запрос.
roll based
````

````sql
SELECT SQL_CACHE
SELECT SQL_NO_CACHE
SELECT SQL_CALC_FOUND_ROWS * FROM table;
SELECT FOUND_ROWS();

SHOW COLUMNS FROM table LIKE '%';

INSERT INTO tbl1 (field1) SELECT field2 FROM tbl2;

SELECT AUTO_INCREMENT FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'databaseName' AND TABLE_NAME = 'table';

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
