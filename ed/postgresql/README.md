PostgreSQL
-

PostgreSQL 10.0

(ORDBMS) Object-Relational Database Management System.

[online config](http://pgtune.leopard.in.ua/).

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

#### Data Types:

boolean

character values:
  char: holds a single character
  char (#): holds # number of characters.
  varchar (#)

integer values:
  smallint
  int
  serial

floating-point values:
  float (#)
  real: 8-byte floating point number
  numeric (#,after_dec)

date and time values:
  date
  time
  timestamp
  timestamptz
  interval

geometric data:
  point
  line
  lseg
  box
  polygon

device specifications:
  inet
  macaddr

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

@TODO:
json field
array field
window functions
ANY/SOME/ALL
parallel query
index diff unique/primary
https://app.pluralsight.com/library/courses/postgresql-advanced-sql-queries/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-sql-queries-introduction/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-advanced-server-programming/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-time-temporal-data/table-of-contents
