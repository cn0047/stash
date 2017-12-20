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
  id serial NOT NULL PRIMARY KEY,
  n int,
  d double precision,
  s character varying(20)
);
create table products (
  product_no integer,
  name text,
  price numeric,
  CHECK (price > 0),
  discounted_price numeric,
  CHECK (discounted_price > 0),
  CHECK (price > discounted_price) 
);
create table example (
  a integer,
  b integer,
  c integer,
  PRIMARY KEY (a, b),
  UNIQUE (a, c) 
);
````

Random stuff:

````
SELECT ARRAY[1, 2, 1+2];
  array
---------
 {1,2,3}

 SELECT ARRAY[1, 2, 1+2, '4', true]::integer[];
    array
-------------
 {1,2,3,4,1}

SELECT ROW(1,2.5,'this is a test');
````

#### JSON:

````
row_to_json(fieldFromDB)
to_json(string)
json_agg(fieldFromDB)
json_object('{a, 1, b, "def", c, 3.5}')
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

@TODO:
https://app.pluralsight.com/library/courses/postgresql-index-tuning-performance-optimization/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-advanced-sql-queries/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-sql-queries-introduction/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-advanced-server-programming/table-of-contents
https://app.pluralsight.com/library/courses/postgresql-time-temporal-data/table-of-contents
