PostgreSQL
-

(ORDBMS) Object-Relational Database Management System.

````
\l
\d
\dt *.*
````

````
pg_dump -h localhost -p 5432 -U dbu -d test > /var/lib/postgresql/data/dump.sql
psql -h localhost -p 5432 -U dbu -d td < /var/lib/postgresql/data/dump.sql
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
