MySql
-
<br>5.7.14-google-log
<br>5.1.72-rel14.10

#### Snippets

````sql
-- Create new user like root
-- mysql --user=root mysql
CREATE USER 'test2'@'localhost' IDENTIFIED BY 'pass';
GRANT ALL PRIVILEGES ON *.* TO 'test2'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON testDb.* TO 'client'@'%' IDENTIFIED BY 'pwd' WITH GRANT OPTION

-- grant user
GRANT ALL PRIVILEGES ON testDB.* TO 'test2'@'localhost' IDENTIFIED BY 'pass';
GRANT ALL PRIVILEGES ON testDB.* TO 'test2'@'%' IDENTIFIED BY 'pass';

-- set password for user
SET PASSWORD FOR 'ziipr'@'localhost' = PASSWORD('12345');
DROP USER 'jeffrey'@'localhost';
````

````sql
CREATE TABLE tmpTbl (
    id int UNSIGNED NOT NULL AUTO_INCREMENT KEY
    , tid TINYINT UNSIGNED NOT NULL DEFAULT 0
    , userId INT UNSIGNED NOT NULL DEFAULT 0
    , message TEXT, -- NOT NULL DEFAULT ''
    , createdAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    , name VARCHAR(25) NOT NULL DEFAULT '' UNIQUE
);

CREATE TABLE IF NOT EXISTS tbl LIKE talbeLike;
INSERT INTO table SELECT * FROM talbeLike;

ALTER  TABLE table    ADD field INTEGER(1) NOT NULL AFTER anotherFieldName;
ALTER  TABLE table    ADD UNIQUE KEY (field);
ALTER  TABLE engineer ADD UNIQUE KEY username_email (username, email);
ALTER  TABLE engineer ADD FOREIGN KEY (LanguageID) REFERENCES languages(LanguageId) ON DELETE RESTRICT;
CREATE INDEX indexName ON table (column);

ALTER  TABLE table MODIFY field VARCHAR(7) NOT NULL DEFAULT '';
ALTER  TABLE table CHANGE field field VARCHAR(7) NOT NULL DEFAULT ''; -- rename column
ALTER  TABLE table CHANGE field field TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER  TABLE table DROP COLUMN field;
ALTER  TABLE table DROP FOREIGN KEY key; -- CONSTRAINT.
ALTER  TABLE table DROP KEY key;

DROP   TABLE IF EXISTS tableNewName;
DROP   INDEX indexName ON table;

RENAME TABLE table TO tableNewName;
````

````sql
SHOW ENGINE INNODB STATUS\G -- show lot of helpful info including last error

SHOW OPEN TABLES from  dataBaseName; -- Show locked tables.

SHOW TABLE STATUS WHERE name = 'table'; -- Info about table, with creation date.
SELECT
    table_name, create_time
FROM information_schema.tables
WHERE table_schema='dataBaseName' AND create_time BETWEEN '2014-05-05' AND '2014-05-07'
;
SHOW COLUMNS FROM table LIKE '%'; -- columns in table format like DESC tableName
-- get AUTO_INCREMENT
SELECT AUTO_INCREMENT FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'databaseName' AND TABLE_NAME = 'table';

SELECT SQL_CALC_FOUND_ROWS * FROM table;
SELECT FOUND_ROWS();

SELECT SUM(field) FROM table GROUP BY field WITH ROLLUP;
INSERT LOW_PRIORITY INTO table1 SELECT field FROM table2;

SELECT * FROM stockTable ORDER BY field(stockid, 33, 12, 53); -- ordered result: 33, 12, 53
````

#### Functions

````sql
CONV(2, 10, 2) -- CONV(num , from_base, to_base)
LPAD('ZipCode', 10, '0') -- 000ZipCode

ROW_COUNT()                        -- after insert, update, delete
LAST_INSERT_ID()
CONNECTION_ID()

UNIX_TIMESTAMP()
CURTIME()
sysdate()                          -- now
LAST_DAY(date)                     -- last day in month
CURRENT_DATE                       -- today

REPLACE('vvv.site.com', 'v', 'w')  -- www.site.com

SELECT SUBSTRING_INDEX('www.mysql.com', '.', 1); -- www
SELECT SUBSTRING_INDEX('www.mysql.com', '.', 2); -- www.mysql 
SELECT SUBSTRING_INDEX('www.mysql.com', '.', -1); -- com
SELECT SUBSTRING_INDEX('www.mysql.com', '.', -2); -- mysql.com 
SELECT SUBSTRING_INDEX('yoyo3450@hotmail.com', '@', -1); -- hotmail.com
SELECT SUBSTRING_INDEX(SUBSTRING_INDEX(jsonField, '"name":"', -1), '"', 1) name from users; -- name field from JSON field
SELECT ELT(1, 'foo', 'bar');       -- foo
SELECT FIELD('foo', 'foo', 'bar'); -- 1
SELECT FIND_IN_SET('b', 'a,b,cd'); -- 2
SELECT LEAST(15,10,25);            -- 10

INET_ATON(ip)
INET_NTOA(i)

UUID()
````

Privilege needed to create a function:

````
+----------------+-----------+----------------------------------+
| Privilege      | Context   | Comment                          |
+----------------+-----------+----------------------------------+
| Create routine | Databases | To use CREATE FUNCTION/PROCEDURE |
+----------------+-----------+----------------------------------+
````

A routine is considered DETERMINISTIC if it always produces the same result for the same input parameters.

#### Flashback

````sql
ndb_config # Extract MySQL Cluster Configuration Information.

apt-get install percona-toolkit # tools for performance analyzing
apt-get install percona-xtrabackup
pt-online-schema-change # ALTER tables without locking them
apt-get install sysbench # tool to test mysql performance
check-unused-keys # tool to interact with INDEX_STATISTICS

INSERT INTO brand2 (name) SELECT name FROM brand;

LIMIT OFFSET, COUNT

-- First works AND than OR

LIKE '[JM]%'                          -- begins on J or M
% - *
_ - 1 char
[] - group of symblols

COLLATE UTF8_GENERAL_CI LIKE

SELECT COALESCE(field, 0) FROM table; -- if field is null returns 0

OREDER BY ASC                         -- default ASC

INSERT INTO table VALUES (default);   -- default value

UPDATE tab SET field = DEFAULT(field);

CHECK TABLE tab;    -- checks a table for errors, for MyISAM - update key statistics
ANALYZE TABLE tab;  -- performs a key distribution analysis
REPAIR TABLE tab;   -- repairs a possibly corrupted table (MyISAM, ARCHIVE, and CSV)
OPTIMIZE TABLE tab; -- (defragmentation) reorganizes the physical storage of table data and associated index data,
                    -- to reduce storage space and improve I/O efficiency when accessing the table.

SELECT VERSION();
SELECT USER();
SELECT * FROM mysql.user; -- Select all mysql users.
SHOW GRANTS FOR 'usr';
SHOW GRANTS;
SHOW PRIVILEGES;

SHOW FULL TABLES;

SELECT * INTO OUTFILE '/tmp/users.txt'
FIELDS TERMINATED BY ',' OPTIONALLY ENCLOSED BY '"'
LINES TERMINATED BY '\n'
FROM users;

SHOW PROCESSLIST; -- to see connections

Horizontal scaling means that you scale by adding more machines into your pool of resources (Replication).
Vertical scaling means that you scale by adding more power (CPU, RAM) to your existing machine.
````

#### Truncate Vs delete

TRUNCATE:
It requires the DROP privilege.
Does not invoke ON DELETE triggers.
It cannot be performed for InnoDB tables with parent-child foreign key relationships.
Truncate operations drop and re-create the table.
Cannot be rolled back.
Any AUTO_INCREMENT value is reset to its start value.

DELETE:
The DELETE statement deletes rows from tbl_name and returns the number of deleted rows.
Need the DELETE privilege on a table to delete rows from it.
You cannot delete from a table and select from the same table in a subquery.

#### Tricks

````
~/.mysql_history

tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
mysql> \s
````

#### Set

````sql
SET @id = 1;
SELECT @id;
+------+
| @id  |
+------+
|    1 |
+------+

select @v := (select 100) a, 200 b;

SELECT @myRight := rgt FROM nested_category WHERE name = 'TELEVISIONS';
SELECT @myRight;
+----------+
| @myRight |
+----------+
|        9 |
+----------+
````

#### NULL

Aggregate (summary) functions such as COUNT(), MIN(), and SUM() ignore NULL values.
The exception to this is COUNT(*), which counts rows

Not NULL is important because we do not have to work on code level with int number `0` or `11` or `NULL` (wtf)!
Also all data in DB have CONSISTENT representation.
But `text` datatype is exception((( It's impossible to specify default value for `text`.

It also depends on data stored in table,
if we grab some data from somewhere and we cannot guarantee consistency on app level - have to use NULL.
For example: we have spy bot, and bot can grab only email, or email and phone, or only name...
So it is no way to figure out did bot found empty info or didn't find info,
unless we use NULL.

#### Options

````sql
SET SQL_BIG_SELECTS  = 1;
SET SQL_SAFE_UPDATES = 0;
SET FOREIGN_KEY_CHECKS = 0;

SELECT SQL_CACHE
SELECT SQL_NO_CACHE
````

#### Sizes

Every table (regardless of storage engine) has a maximum row size of 65,535 bytes.

Maximum Names Length:
| Identifier                                                                                                 | Maximum Length (characters) |
|------------------------------------------------------------------------------------------------------------|-----------------------------|
| Database, Table                                                                                            | 64 (NDB engine: 63)         |
| Column, Index, Constraint, Stored Program, View, Tablespace, Server, Log File Group, User-Defined Variable | 64                          |
| Alias                                                                                                      | 256                         |

````sql
-- tables sizes
SELECT
    table_name AS 'Table',
    round(((data_length + index_length) / 1024 / 1024), 2) 'Size in MB'
FROM information_schema.TABLES
WHERE table_schema = 'dbName' AND table_name = 'tableName'
;

-- data usage
SET @tbl1 = 'user';
SELECT
    TABLE_NAME, TABLE_ROWS, AVG_ROW_LENGTH, DATA_LENGTH, INDEX_LENGTH, DATA_FREE
FROM information_schema.TABLES
WHERE table_schema = 'test' AND table_name IN (@tbl1)
;

-- Count of rows in db.
SELECT TABLE_NAME, TABLE_ROWS FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'dbName';
````

#### mySqlDump

````sh
mysqldump -h hostname -u user -pPassWord --skip-triggers --single-transaction --complete-insert --extended-insert --quote-names --disable-keys dataBaseName | gzip -с > DUMP_dataBaseName.sql.gz
# dump only db schema
mysqldump -h hostname -u user -pPassWord --no-data --single-transaction --complete-insert --extended-insert --quote-names dbname table > dump.sql.gz
-- mysqldump -hHost Base table | gzip | uuencode table.gz | mail mail@com.com -s table
-- mysqldump -h Host Base table --where="id=11" | mail mail@com.com
mysql -hHost -uUser -pPass -DBase < dumpFile.sql
````
````sh
-- mysql -h Host -D Base -e "select * from table where id in (1,2);" | gzip | pv | uuencode result.csv.gz | mail
````

#### Slow query log

````sh
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

#### Conf

/etc/mysql/my.cnf
````sql
slow_query_log = 1
slow_query_log_file = /var/log/mysql/logSlowQueries.log
long_query_time = 1
log_queries_not_using_indexes = 0
innodb_log_file_size = 5M

# save from memory to disc:
# O_DIRECT - 4 direct writes on disc;
# O_DSYNC - 2 direct writes & 2 async;
innodb_flush_method

# save to disc after transaction (strategy how to write data on disc):
innodb_flush_log_at_trx_commit

# redo log (log of actions in case of corruption) size, mysql will do all this actions after recover
innodb_log_file_size

# memory for reading tables, indexes, etc
innodb_buffer_pool_size = 7G # 80% from OS memory

# buffer for uncommitted transactions
innodb_log_buffer_size

innodb_file_per_table = ON

query_cache_limit = 1M
query_cache_size  = 8M
# mysql> SHOW STATUS LIKE '%Qcache%';
# mysql> FLUSH QUERY CACHE

# threads for new client's connections
thread_cache_size = 128
# mysql> SHOW GLOBAL STATUS LIKE 'Threads_created';

max_connections = 256

# for connection timeout
wait_timeout = 600

# lock for concurrent transactions
innodb_lock_wait_timeout=100

innodb_force_recovery = 1 # 1, 2, 3, ...

# ignore tables for replication
replicate_wild_ignore_table = playgroundDB.%
replicate_wild_ignore_table = dataDB.temp_%

# don't save binlog to disk for better performance
sync_binlog = 0
# slave threads count
slave-parallel-workers = 2
slave-parallel-type = LOGICAL_CLOCK

# for fast mysql reload
innodb_buffer_pool_dump_at_shutdown = ON
innodb_buffer_pool_load_at_startup = ON

general-log = 1
general-log-file = /var/log/mysql/general.log
# And run:
# sudo mkdir -p /var/log/mysql
# sudo touch /var/log/mysql/general.log
# sudo chown mysql:mysql /var/log/mysql/general.log

SET global general_log_file='/var/log/mysql/general.log';
SET global general_log_file='/tmp/mysql.general.log';
SET global general_log = 1;

SET profiling = 1;
SHOW PROFILES;
SHOW PROFILE;
SHOW PROFILE FOR QUERY 1;
SHOW PROFILE CPU FOR QUERY 2; -- ALL, BLOCK IO, CONTEXT SWITCHES, CPU, IPC, PAGE FAULTS, SOURCE, SWAPS
````

#### Storage engines

* InnoDB
    * Support for **transactions** (giving you support for the ACID property).
    * **Foreign key** constraints (take more time in designing).
    * **Row level locking**. Faster in **write-intensive** because it utilizes row-level locking and only hold up changes to the same row that’s being inserted or updated.
    * **Recovers from a crash** or other unexpected shutdown by replaying its logs.
    * Consumes more system resources such as RAM.
* MyISAM
    * **Table level locking**. Slower than InnoDB for tables that are frequently being inserted to or updated, because the entire table is locked for any insert or update.
    * Faster than InnoDB on the whole as a result of the simpler structure thus much less costs of server resources.
    * Especially good for **read-intensive** (select) tables.
    * (The maximum number of indexes per MyISAM table is 64. The maximum number of columns per index is 16).
    * (Uses one file for data rows and another for index records).
* MEMORY
* CSV
* ARCHIVE
* BLACKHOLE
* MERGE
* FEDERATED
* EXAMPLE
* NDB

#### [Data Types](http://dev.mysql.com/doc/refman/5.0/en/data-types.html)

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
    * Enum (You cannot employ a user variable as an enumeration value, an enumeration value can also be the empty string ('') or NULL).
    * Set (Is a string object that can have zero or more values, can have a maximum of 64 distinct members, cannot contain comma).
* Spatial Data Types:
    * Geometry.
    * Point.
    * Linestring.
    * Polygon.
    * Multipoint.
    * Multilinestring.
    * Multipolygon.
    * Geometrycollection.
