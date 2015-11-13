MySql
-

*5.1.72-rel14.10*

####Snippets
````sql
INSERT INTO brand2 (name) SELECT name FROM brand;

ALTER  TABLE table ADD field INTEGER(1) NOT NULL;
ALTER  TABLE table ADD UNIQUE KEY (field);
ALTER  TABLE engineer ADD UNIQUE KEY username_email (username, email);
ALTER  TABLE table MODIFY field VARCHAR(7) NOT NULL DEFAULT '';
ALTER  TABLE table CHANGE field field VARCHAR(7) NOT NULL DEFAULT ''; -- rename column
ALTER  TABLE table CHANGE field field TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER  TABLE table DROP COLUMN field;
ALTER  TABLE table DROP FOREIGN KEY key; -- CONSTRAINT.
ALTER  TABLE table DROP KEY key;
CREATE TABLE IF NOT EXISTS table LIKE talbeLike; INSERT INTO table SELECT * FROM talbeLike;
RENAME TABLE table TO tableNewName;
CREATE INDEX indexName ON table (column);
DROP   TABLE IF EXISTS tableNewName;
DROP   INDEX indexName ON table;

SHOW OPEN TABLES from  dataBaseName; -- Show locked tables.

SHOW TABLE STATUS WHERE name = 'table'; -- Info about table, with creation date.
SELECT
    table_name, create_time
FROM information_schema.tables
WHERE table_schema='dataBaseName' AND create_time BETWEEN '2014-05-05' AND '2014-05-07'
;
SHOW COLUMNS FROM table LIKE '%';
SELECT AUTO_INCREMENT FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'databaseName' AND TABLE_NAME = 'table';

SELECT SQL_CALC_FOUND_ROWS * FROM table;
SELECT FOUND_ROWS();

SELECT SUM(field) FROM table GROUP BY field WITH ROLLUP;
INSERT LOW_PRIORITY INTO table1 SELECT field FROM table2;

SELECT * FROM stockTable ORDER BY field(stockid, 33, 12, 53); -- ordered result: 33, 12, 53
````

####NULL
Aggregate (summary) functions such as COUNT(), MIN(), and SUM() ignore NULL values.
The exception to this is COUNT(*), which counts rows

####Optimizations
````sql
force index (createdAt)

// SELECT DISTINCT faster than SELECT GROUB BY.
SELECT DISTINCT id FROM table;
// vs
SELECT id FROM table GROUP BY id;
````

####Options
````sql
SET SQL_BIG_SELECTS  = 1;
SET SQL_SAFE_UPDATES = 0;

SELECT SQL_CACHE
SELECT SQL_NO_CACHE
````

####Sizes
````sql
SELECT
    table_name AS 'Table',
    round(((data_length + index_length) / 1024 / 1024), 2) 'Size in MB'
FROM information_schema.TABLES
WHERE table_schema = 'dbName' AND table_name = 'tableName'
;

SELECT
    TABLE_NAME, TABLE_ROWS, AVG_ROW_LENGTH, DATA_LENGTH, INDEX_LENGTH, DATA_FREE
FROM information_schema.TABLES
WHERE table_schema = 'test' AND table_name IN ('testBinUniqueKey20', 'testBinUniqueKey200')
;

SELECT
    CONCAT(table_schema, '.', table_name),
    CONCAT(ROUND(table_rows / 1000000, 2), 'M')                                    rows,
    CONCAT(ROUND(data_length / ( 1024 * 1024 * 1024 ), 2), 'G')                    data,
    CONCAT(ROUND(index_length / ( 1024 * 1024 * 1024 ), 2), 'G')                   idx,
    CONCAT(ROUND(( data_length + index_length ) / ( 1024 * 1024 * 1024 ), 2), 'G') total_size,
    ROUND(index_length / data_length, 2)                                           idx_frac
FROM information_schema.TABLES
ORDER BY data_length + index_length DESC
LIMIT 50
;

-- Count of rows in db.
SELECT TABLE_NAME, TABLE_ROWS FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'dbName';
````

####Detection of bad index
````sql
SELECT
    t.TABLE_NAME,
    SUM(t.ROWS_READ) AS raw_readed,
    SUM(i.ROWS_READ) AS key_readed,
    ROUND((SUM(i.ROWS_READ)/SUM(t.ROWS_READ))*100, 2) AS index_coverage
FROM information_schema.TABLE_STATISTICS t
LEFT join information_schema.INDEX_STATISTICS i ON t.TABLE_SCHEMA = i.TABLE_SCHEMA AND t.TABLE_NAME = i.TABLE_NAME
WHERE t.TABLE_SCHEMA = 'dbName'
GROUP BY t.TABLE_NAME
HAVING raw_readed > 10000
ORDER BY raw_readed DESC
;
````

####Functions
````sql
CONV(2, 10, 2) -- CONV(num , from_base, to_base)
LPAD(ZipCode, 5, '0')

ROW_COUNT()                        -- after insert, update, delete
LAST_INSERT_ID()
CONNECTION_ID()

UNIX_TIMESTAMP()
CURTIME()
sysdate()                          -- now
LAST_DAY(date)                     -- last day in month
CURRENT_DATE                       -- today

REPLACE('vvv.site.com', 'v', 'w')

SELECT ELT(1, 'foo', 'bar');       -- foo
SELECT FIELD('foo', 'foo', 'bar'); -- 1
SELECT FIND_IN_SET('b', 'a,b,cd'); -- 2
SELECT LEAST(15,10,25);            -- 10

INET_ATON(ip)
INET_NTOA(i)

UUID()
````

####mySqlDump
````
mysqldump -h hostname -u user -pPassWord --skip-triggers --single-transaction --complete-insert --extended-insert --quote-names --disable-keys dataBaseName | gzip -с > DUMP_dataBaseName.sql.gz
-- mysqldump -h hostname -u user -d --skip-triggers --single-transaction --complete-insert --extended-insert --quote-names dbname table |gzip > sql.gz
-- mysqldump -hHost Base table | gzip | uuencode table.gz | mail mail@com.com -s table
-- mysqldump -h Host Base table --where="id=11" | mail mail@com.com
-- php-mysqldump -h HOST -u USER -a table -f sql -e "SELECT * FROM table LIMIT 10"
 mail@com.com -s "query result"
-- zcat table.gz | host -uUser -pPass base
-- zcat DUMP_dataBaseName.sql.gz | mysql
mysql -hHost -uUser -pPass -DBase < dumpFile.sql
````
````
-- mysql -h Host -D Base -te "SELECT NOW()" | mail mail@com.com
-- mysql -hHost -uUser -pPass base ~ table
-- mysql -h Host -D Base -e "select * from table where id in (1,2);" | gzip | pv | uuencode result.csv.gz | mail
````

####Flashback
````sql
LIMIT OFFSET, COUNT

-- First works AND than OR

LIKE '[JM]%'                          -- begins on J or M
% - *
_ - 1 char
[] - group of symblols

COLLATE UTF8_GENERAL_CI LIKE

SELECT COALESCE(field, 0) FROM table; -- if field is null returns 0

OREDER BY ASC                         -- default ASC

LEFT OUTER JOIN                       -- search rows that not exists in second table

INSERT INTO table VALUES (default);   -- default value

UPDATE tab SET field = DEFAULT(field);

CHECK TABLE tab;
ANALYZE TABLE tab;
REPAIR TABLE tab;
OPTIMIZE TABLE tab;

SELECT VERSION();
SELECT USER();
SELECT * FROM mysql.user; -- Select all mysql users.
SHOW GRANTS FOR 'usr';
SHOW GRANTS;
SHOW PRIVILEGES;

SHOW FULL TABLES;

Horizontal scaling means that you scale by adding more machines into your pool of resources (Replication).
Vertical scaling means that you scale by adding more power (CPU, RAM) to your existing machine.
````

####Slow query log
````
mkdir /var/log/mysql
touch /var/log/mysql/logSlowQueries.log
chown mysql.mysql -R /var/log/mysql

mysql -e "
SET GLOBAL slow_query_log_file = '/var/log/mysql/logSlowQueries.log';
SET GLOBAL slow_query_log = 1;
SET GLOBAL long_query_time = 1;
SET GLOBAL log_queries_not_using_indexes = 0;
"
````

Conf:
/etc/mysql/my.cnf
````sql
slow_query_log = 1
slow_query_log_file = /var/log/mysql/logSlowQueries.log
long_query_time = 1
log_queries_not_using_indexes = 0
innodb_log_file_size = 5M
````
````sql
SET global general_log_file='/tmp/myslq.log';
SET global general_log = 0;
````

####Tricks
````
~/.mysql_history

tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
mysql> \s
````

####Storage engines
* InnoDB
    * Support for transactions (giving you support for the ACID property).
    * Foreign key constraints (take more time in designing).
    * Row level locking. Faster in write-intensive because it utilizes row-level locking and only hold up changes to the same row that’s being inserted or updated.
    * Recovers from a crash or other unexpected shutdown by replaying its logs.
    * Consumes more system resources such as RAM.
* MyISAM
    * Table level locking. Slower than InnoDB for tables that are frequently being inserted to or updated, because the entire table is locked for any insert or update.
    * Faster than InnoDB on the whole as a result of the simpler structure thus much less costs of server resources.
    * Especially good for read-intensive (select) tables.
* MEMORY
* CSV
* ARCHIVE
* BLACKHOLE
* MERGE
* FEDERATED
* EXAMPLE

####[Data Types](http://dev.mysql.com/doc/refman/5.0/en/data-types.html)
* Numeric Types:
    * [Integer](http://dev.mysql.com/doc/refman/5.0/en/integer-types.html):
        * Tinyint (Bytes 1).
        * Smallint (Bytes 2).
        * Mediumint (Bytes 3).
        * Int (Bytes 4).
        * Bigint (Bytes 8).
    * Fixed-Point:
        * Decimal.
        * Numeric.
    * Floating-Point:
        * Float (A precision from 0 to 23 results in a 4-byte single-precision FLOAT column.),
        * Double (A precision from 24 to 53 results in an 8-byte double-precision DOUBLE column.),
    * Bit-Value:
        * Bit.
* Date and Time Types:
    * Date (The supported range is '1000-01-01' to '9999-12-31'.).
    * Datetime (The supported range is '1000-01-01 00:00:00' to '9999-12-31 23:59:59'.).
    * Timestamp (Has a range of '1970-01-01 00:00:01' UTC to '2038-01-19 03:14:07' UTC.).
    * Time (Range from '-838:59:59' to '838:59:59'.).
    * Year (Range of 1901 to 2155, or 0000.).
    * YEAR(2) (Range of 1901 to 2155, and 0000) (Display only the last two digits, omit the century digits).
    * YEAR(4) (Range of 1901 to 2155, and 0000).
* String Types:
    * Char.
    * Varchar.
    * Binary.
    * Varbinary.
    * Blob:
        * Tinyblob (2^8-1 bytes).
        * Blob (2^16-1 bytes).
        * Mediumblob (2^24-1 bytes).
        * Longblob (2^32-1 bytes).
    * Text:
        * Tinytext (255 bytes).
        * Text (65535 bytes).
        * Mediumtext (16 777 215 bytes).
        * Longtext (4 294 967 295 bytes).
    * Enum.
    * Set (Is a string object that can have zero or more values).
* Spatial Data Types:
    * Geometry.
    * Point.
    * Linestring.
    * Polygon.
    * Multipoint.
    * Multilinestring.
    * Multipolygon.
    * Geometrycollection.
