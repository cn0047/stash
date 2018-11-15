PostgreSQL
-
<br>10.0

[online config](http://pgtune.leopard.in.ua/).

(ORDBMS) Object-Relational Database Management System.
`pgfouine` - profiler.

DB can contains schema, and schema contains tables.
Schema - like directory in file system.

````bash
# init on osx
brew services start postgresql
createdb cws
psql -d cws -L /tmp/p.log
CREATE USER usr WITH PASSWORD 'pass';
GRANT ALL PRIVILEGES ON database dbName TO usr;

psql -d postgres://dbu:dbp@localhost/test
psql -h localhost -U dbu -d test -c 'select 204'

pg_dump -h localhost -p 5432 -U dbu -d test > /var/lib/postgresql/data/dump.sql
psql -h localhost -p 5432 -U dbu -d td < /var/lib/postgresql/data/dump.sql

pg_ctl status

# general log
# SELECT current_setting('log_statement');
# SET log_statement='all';
# ALTER DATABASE dbname SET log_statement='all';

pg_trgm # Trigram (Trigraph) module
````

````sql
set x '''1''';
select :x;

set x.y = 204;
select current_setting('x.y');
````

#### REPL:

````
\pset pager on
\pset pager always
\pset pager off
-- \setenv PAGER less

# show databases
\l
\d
# use db
\c unittests

# show tables
\dt *.*
\dt *test*
# more info
\d+ viewName

# `SHOW CREATE TABLE`
pg_dump -t tableName --schema-only
````

````sql
COPY tableName TO 'filePath' CSV (DELIMER ',');
COPY tableName FROM 'filePath' DELIMER ',';
````

````
VACUUM VERBOSE tableName; -- helpful info
````

#### System Columns:

* `oid` - Object identifier (table was created using WITH OIDS).
* `tableoid` - OID of the table containing this row.
* `xmin` - The identity (transaction ID) of the inserting transaction for this row version.
* `cmin` - Command identifier within the inserting transaction.
* `xmax` - The identity (transaction ID) of the deleting transaction.
* `cmax` - The command identifier within the deleting transaction.
* `ctid` - The physical location of the row version within its table.

#### Rules

* INSTEAD
* ALSO (additional command to original command)

Triggers better than rules.

#### PL - Procedural Languages.

* PL/pgSQL (default and has optimal performance)
* PL/Tcl
* PL/Perl
* PL/Python

We can return multiple values from procedures.

#### Window Functions

Like aggregate functions but doesn't cause rows to become grouped into a single output row.
it is valid to include an aggregate function call in the arguments of a window function,
but not vice versa.

Example:

````
SELECT depname, empno, salary, avg(salary) OVER (PARTITION BY depname)
FROM empsalary;
````

#### Parallel Query

Only `Seq Scan` can be parallelized,
but indexed column disable parallelism.

````
SHOW max_parallel_workers_per_gather;
SET max_parallel_workers_per_gather = 4;
````

Also `dynamic_shared_memory_type` must be set to a value other than none,
And the system must NOT be running in single-user mode.

#### Configuration:

`/etc/postgresql/8.3/main/postgresql.conf`
`/usr/local/var/postgres/postgresql.conf`

````
listen_addresses = '*'
max_connections
shared_buffers # cache, must be 15-25% from OS memory
effective_cache_size # memory for disk cache, must be 50-75% from OS memory

# for intensive writes
checkpoint_segments
wal_buffers
synchronous_commit

work_mem
maintainance_work_mem

log_destination = 'syslog'
redirect_stderr = off
silent_mode = on
syslog_facility = 'LOCAL0'
syslog_ident = 'postgres'

log_min_duration_statement = 0
log_duration = on
log_statement = 'none'
````

Dev:

````
# logs
log_statement = 'all'
logging_collector = on
log_min_duration_statement = 0
log_connections = on
# target
log_destination = 'csvlog'
log_directory = '/tmp'
log_filename = 'psql.log'
# verbosity
client_min_messages = notice
log_min_messages = info
log_min_error_statement = info
# output
debug_pretty_print = on
debug_print_parse = off
debug_print_plan = off
debug_print_rewritten = off
````

#### Data Types:

boolean

* character values:
  * char: holds a single character
  * char (#): holds # number of characters.
  * varchar (#)

* integer values:
  * smallint
  * int
  * serial

* floating-point values:
  * float (#)
  * real: 8-byte floating point number
  * numeric (#,after_dec)

* date and time values:
  * date
  * time
  * timestamp
  * timestamptz
  * interval

* geometric data:
  * point
  * line
  * lseg
  * box
  * polygon

* device specifications:
  * inet
  * macaddr

Pseudo-Types:

* any
* anyelement
* anyarray
* anynonarray
* anyenum
* anyrange
* cstring
* internal
* language_handler
* fdw_handler
* index_am_handler
* tsm_handler
* record
* trigger
* event_trigger
* pg_ddl_command
* void
* unknown
* opaque
