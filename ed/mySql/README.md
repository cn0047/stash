MySql
-

*5.1.72-rel14.10*

####Snippets
````sql
CREATE TABLE tableName (
KEY `accounting_user` (`accountingid`,`user_id`)
);
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

####INDEX
If Cardinality to low - index will not uses!

####NULL
Aggregate (summary) functions such as COUNT(), MIN(), and SUM() ignore NULL values.
The exception to this is COUNT(*), which counts rows

####Optimizations
````sql
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

####Tricks
````
~/.mysql_history

tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
mysql> \s
````

####Lock and unlock tables
Locks may be used to emulate transactions or to get more speed when updating tables.
A session holding a WRITE lock can perform table-level operations such as DROP TABLE or TRUNCATE TABLE.
For sessions holding a READ lock, DROP TABLE and TRUNCATE TABLE operations are not permitted.
LOCK TABLES is permitted (but ignored) for a TEMPORARY table.

Rules for Lock Acquisition:
* READ [LOCAL] lock:
    * The session that holds the lock can read the table (but not write it).
    * Multiple sessions can acquire a READ lock for the table at the same time.
    * Other sessions can read the table without explicitly acquiring a READ lock.
    * The LOCAL modifier enables nonconflicting INSERT statements (concurrent inserts) by other sessions
      to execute while the lock is held.
* [LOW_PRIORITY] WRITE lock:
    * The session that holds the lock can read and write the table.
    * Only the session that holds the lock can access the table.
      No other session can access it until the lock is released.
    * Lock requests for the table by other sessions block while the WRITE lock is held.

WRITE locks normally have higher priority than READ locks.

````sql
LOCK TABLES t1 READ;
SELECT COUNT(*) FROM t1;
+----------+
| COUNT(*) |
+----------+
|        3 |
+----------+
SELECT COUNT(*) FROM t2;
ERROR 1100 (HY000): Table 't2' was not locked with LOCK TABLES
````

You cannot refer to a locked table multiple times in a single query using the same name.
Use aliases instead, and obtain a separate lock for the table and each alias:
````sql
LOCK TABLE t WRITE, t AS t1 READ;
INSERT INTO t SELECT * FROM t;
ERROR 1100: Table 't' was not locked with LOCK TABLES
INSERT INTO t SELECT * FROM t AS t1;
````

Rules for Lock Release:
* A session can release its locks explicitly with UNLOCK TABLES.
* If a session issues a LOCK TABLES statement to acquire a lock while already holding locks,
  its existing locks are released implicitly before the new locks are granted.
* If a session begins a transaction (for example, with START TRANSACTION),
  an implicit UNLOCK TABLES is performed, which causes existing locks to be released.
* If the connection for a client session terminates, whether normally or abnormally,
  the server implicitly releases all table locks held by the session.
* If you use ALTER TABLE on a locked table, it may become unlocked.

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
