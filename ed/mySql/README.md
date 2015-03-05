MySql
-

*5.1.72-rel14.10*



####Snippets
````sql
ALTER  TABLE table ADD field INTEGER(1) NOT NULL;
ALTER  TABLE table ADD UNIQUE KEY (field);
ALTER  TABLE table MODIFY field VARCHAR(7) NOT NULL DEFAULT '';
ALTER  TABLE table CHANGE field field VARCHAR(7) NOT NULL DEFAULT ''; -- rename column
ALTER  TABLE table CHANGE field field TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;
ALTER  TABLE table DROP COLUMN field;
ALTER  TABLE table DROP FOREIGN KEY key; -- CONSTRAINT.
ALTER  TABLE table DROP KEY key;
CREATE TABLE table LIKE talbeLike; INSERT INTO table SELECT * FROM talbeLike;
RENAME TABLE table TO tableNewName;
CREATE INDEX indexName ON table (column);
DROP   INDEX indexName ON table;

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
WHERE table_schema = 'dbName' AND table_name = 'tableName';
````

####Functions
````sql
ROW_COUNT()                        -- after insert, update, delete
LAST_INSERT_ID()
CONNECTION_ID()

UNIX_TIMESTAMP()
CURTIME()
sysdate()                          -- now
LAST_DAY(date)                     -- last day in month

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
mysqldump -h hostname -u user --skip-triggers --single-transaction --complete-insert --extended-insert --quote-names --disable-keys dataBaseName | gzip -с > DUMP_dataBaseName.sql.gz
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
SHOW GRANTS FOR 'usr';
SHOW GRANTS;
SHOW PRIVILEGES;

SHOW FULL TABLES;

Horizontal scaling means that you scale by adding more machines into your pool of resources (Replication).
Vertical scaling means that you scale by adding more power (CPU, RAM) to your existing machine.
````

####Tricks
````
~/.mysql_history

tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
````

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

##History of MySQL

####MySQL 5.7
* Security improvements.
* SQL mode changes.
* Online ALTER TABLE (RENAME INDEX).
* InnoDB enhancements.
* Condition handling.
* Optimizer.
* Triggers (Multiple triggers are permitted.).
* Logging.
* Test suite.
* mysql client.
* Database name rewriting with mysqlbinlog.
* HANDLER with partitioned tables.
* Index condition pushdown support for partitioned tables.
* WITHOUT VALIDATION support for ALTER TABLE ... EXCHANGE PARTITION.
* Master dump thread improvements.
* Globalization improvements.
* Changing the replication master without STOP SLAVE.
* Deprecated EXPLAIN EXTENDED and PARTITIONS.
* INSERT DELAYED is no longer supported.

####MySQL 5.6
* Security improvements.
* MySQL Enterprise.
* Changes to server defaults.
* InnoDB enhancements (FULLTEXT indexes, ALTER TABLE without blocking, DATA DIRECTORY clause of the CREATE TABLE (which allows you to create InnoDB file-per-table tablespaces)).
* Partitioning (Maximum number of partitions is increased to 8192).
* Performance Schema.
* MySQL Cluster.
* Replication and logging (Transaction-based replication).
* Optimizer enhancements (EXPLAIN for DELETE, INSERT, REPLACE, UPDATE.).
* Condition handling.
* Data types (Permits fractional seconds for TIME, DATETIME, and TIMESTAMP).
* Host cache.
* OpenGIS.

####MySQL 5.5
* MySQL Enterprise Thread Pool.
* MySQL Enterprise Audit (audit_log).
* Pluggable authentication.
* Multi-core scalability.
* InnoDB I/O subsystem.
* Diagnostic improvements.
* Solaris.
* Default storage engine (Is InnoDB).
* MySQL Cluster.
* Semisynchronous replication.
* Unicode.
* Partitioning (RANGE COLUMNS, LIST COLUMNS, and else).
* SIGNAL and RESIGNAL.
* Metadata locking.
* IPv6 support.
* XML (LOAD XML INFILE).
* Build configuration.

####MySQL 5.1
* Partitioning.
* Row-based replication.
* Plugin API.
* Event scheduler.
* Server log tables.
* Upgrade program.
* MySQL Cluster.
* Backup of tablespaces.
* Improvements to INFORMATION_SCHEMA.
* XML functions with XPath support.
* Load emulator ([mysqlslap](http://dev.mysql.com/doc/refman/5.1/en/mysqlslap.html) program).

####MySQL 5.0
* Information Schema.
* Instance Manager.
* Precision Math.
* Storage Engines.
* Stored Routines.
* Triggers.
* Views.
* Cursors.
* Strict Mode and Standard Error Handling.
* VARCHAR Data Type.
* BIT Data Type.
* Optimizer enhancements.
* XA Transactions.
