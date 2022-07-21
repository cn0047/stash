Administration
-

````sh
ndb_config                         # Extract MySQL Cluster Configuration Information.

apt-get install percona-toolkit    # tools for performance analyzing
apt-get install percona-xtrabackup
pt-online-schema-change            # ALTER tables without locking them
apt-get install sysbench           # tool to test mysql performance
check-unused-keys                  # tool to interact with INDEX_STATISTICS
````

````sh
service mysql status|reload|restart
````

````sql
CHECK TABLE tab;    -- checks a table for errors, for MyISAM - update key statistics
ANALYZE TABLE tab;  -- performs a key distribution analysis
REPAIR TABLE tab;   -- repairs a possibly corrupted table (MyISAM, ARCHIVE, and CSV)
OPTIMIZE TABLE tab; -- (defragmentation) reorganizes the physical storage of table data and associated index data,
                    -- to reduce storage space and improve I/O efficiency when accessing the table.

SHOW ERRORS LIMIT 1;
SHOW WARNINGS LIMIT 1;

SELECT CONNECTION_ID();

SELECT VERSION();
SELECT USER();            -- Select current user.
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

-- Create new user like root
-- mysql --user=root mysql
CREATE USER 'test2'@'localhost' IDENTIFIED BY 'pass';
GRANT ALL PRIVILEGES ON *.* TO 'test2'@'localhost' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON testDb.* TO 'client'@'%' IDENTIFIED BY 'pwd' WITH GRANT OPTION

-- grant user
GRANT ALL PRIVILEGES ON testDB.* TO 'test2'@'localhost' IDENTIFIED BY 'pass';
GRANT ALL PRIVILEGES ON testDB.* TO 'test2'@'%' IDENTIFIED BY 'pass';

-- set password for user
SET PASSWORD FOR 'zii'@'localhost' = PASSWORD('12345');
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
SHOW PROCEDURE STATUS like 'proc';
SHOW CREATE PROCEDURE proc;

SHOW ENGINE INNODB STATUS\G -- show lot of helpful info including last error

SHOW OPEN TABLES from  dataBaseName; -- Show locked tables.

SHOW TABLE STATUS WHERE name = 'table'; -- Info about table: creation date, rows count.
SELECT
    table_name, create_time
FROM information_schema.tables
WHERE table_schema='dataBaseName' AND create_time BETWEEN '2014-05-05' AND '2014-05-07'
;
SHOW COLUMNS FROM table LIKE '%'; -- columns in table format like DESC tableName
-- get AUTO_INCREMENT
SELECT AUTO_INCREMENT FROM information_schema.TABLES
WHERE TABLE_SCHEMA = 'databaseName' AND TABLE_NAME = 'table';
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

#### Tricks

````sh
~/.mysql_history

tee /tmp/out
cat /tmp/out | mail mail@com.com

user@ubuntu:~$ mysql --pager='less -S'

mysql> pager less -SFX
mysql> \s
````

#### Options

````sql
SET SQL_BIG_SELECTS  = 1;
SET SQL_SAFE_UPDATES = 0;
SET FOREIGN_KEY_CHECKS = 0;

SELECT SQL_CACHE
SELECT SQL_NO_CACHE ... -- disable query cache
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

SHOW VARIABLES where Variable_name = 'slow_query_log';
SHOW VARIABLES where Variable_name = 'datadir';
"
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
mysqldump -h hostname -u user -pPassWord --skip-triggers --single-transaction --complete-insert --extended-insert --quote-names --disable-keys dataBaseName \
  | gzip -с > DUMP_dataBaseName.sql.gz
# dump only db schema
mysqldump -h hostname -u user -pPassWord --no-data --single-transaction --complete-insert --extended-insert --quote-names dbname table \
  > dump.sql.gz
-- mysqldump -hHost Base table | gzip | uuencode table.gz | mail mail@com.com -s table
-- mysqldump -h Host Base table --where="id=11" | mail mail@com.com
mysql -hHost -uUser -pPass -DBase < dumpFile.sql
````
````sh
-- mysql -h Host -D Base -e "select * from table where id in (1,2);" | gzip | pv | uuencode result.csv.gz | mail
````

#### Conf

````sql
SET global general_log_file='/var/log/mysql/general.log';
SET global general_log_file='/tmp/mysql.general.log';
SET global general_log = 1;

SET PROFILING = 1; -- enable profiling
SHOW PROFILES;
SHOW PROFILE;
SHOW PROFILE FOR QUERY 1;
SHOW PROFILE CPU FOR QUERY 2; -- ALL, BLOCK IO, CONTEXT SWITCHES, CPU, IPC, PAGE FAULTS, SOURCE, SWAPS
````

/etc/mysql/my.cnf
````sql
[client]
port                        = 3306
socket                      = /var/run/mysqld/mysqld.sock

[mysqld_safe]
socket                      = /var/run/mysqld/mysqld.sock
nice                        = 0

[mysqldump]
quick
quote-names
max_allowed_packet          = 16M

[mysqld]
user                        = mysql
pid-file                    = /var/run/mysqld/mysqld.pid
socket                      = /var/run/mysqld/mysqld.sock
port                        = 3306
basedir                     = /usr
datadir                     = /var/lib/mysql
tmpdir                      = /tmp
language                    = /usr/share/mysql/english
old_passwords               = 0
bind-address                = 127.0.0.1

skip-external-locking

max_allowed_packet          = 16M
key_buffer_size             = 16M

query_cache_size            = 0

expire_logs_days            = 10
max_binlog_size             = 100M

slow_query_log = 1
slow_query_log_file = /var/log/mysql/logSlowQueries.log
long_query_time = 1
log_queries_not_using_indexes = 0
innodb_log_file_size = 5M

innodb_file_per_table       = 1

# save from memory to disc:
# O_DIRECT - 4 direct writes on disc;
# O_DSYNC - 2 direct writes & 2 async;
innodb_flush_method         = O_DIRECT

# save to disc after transaction (strategy how to write data on disc):
innodb_flush_log_at_trx_commit  = 0

# redo log (log of actions in case of corruption) size, mysql will do all this actions after recover
innodb_log_file_size

# memory for reading tables, indexes, etc
innodb_buffer_pool_size = 7G # 80% from OS memory

# 1 gb
innodb_buffer_pool_size     = 512M
max_connections             = 132
# 2 gb
innodb_buffer_pool_size     = 1024M
max_connections             = 136
# 4 gb
innodb_buffer_pool_size     = 2048M
max_connections             = 144
# 8 gb
innodb_buffer_pool_size     = 4096M
max_connections             = 160
# 16 gb
innodb_buffer_pool_size     = 8192M
max_connections             = 192
# 32 gb
innodb_buffer_pool_size     = 16384M
max_connections             = 256
# 64 gb
innodb_buffer_pool_size     = 32768M
max_connections             = 384
# 128 gb
innodb_buffer_pool_size     = 65536M
max_connections             = 640

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
````
