PostgreSQL
-

(ORDBMS) Object-Relational Database Management System.

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
