PostgreSQL
-

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
pg_dump -h localhost -p 5432 -U dbu -d test > /var/lib/postgresql/data/dump.sql
psql -h localhost -p 5432 -U dbu -d td < /var/lib/postgresql/data/dump.sql
````

````sql
create table test (
  id serial not null primary key,
  n int,
  d double precision,
  s character varying(20)
);
````

#### JSON:

````
row_to_json(fieldFromDB)
to_json(string)
json_agg(fieldFromDB)
json_object('{a, 1, b, "def", c, 3.5}')
````

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

@TODO:
inherits
https://app.pluralsight.com/library/courses/postgresql-index-tuning-performance-optimization/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-advanced-sql-queries/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-sql-queries-introduction/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-advanced-server-programming/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-time-temporal-data/table-of-contents
