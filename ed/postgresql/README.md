PostgreSQL
-

PostgreSQL 10.0

(ORDBMS) Object-Relational Database Management System.

[online config](http://pgtune.leopard.in.ua/).

DB can contains schema, and schema contains tables.
Schema - like directory in file system.

#### REPL:

````
\pset pager off

# show databases
\l
\d
# show tables
\dt *.*
\dt *test*
````

````
COPY tableName TO 'filePath' CSV (DELIMER ',');
COPY tableName FROM 'filePath' DELIMER ',';

psql -d postgres://dbu:dbp@localhost/test
psql -h localhost -U dbu -d test -c 'select 204'

pg_dump -h localhost -p 5432 -U dbu -d test > /var/lib/postgresql/data/dump.sql
psql -h localhost -p 5432 -U dbu -d td < /var/lib/postgresql/data/dump.sql
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
