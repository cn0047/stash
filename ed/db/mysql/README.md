MySql
-
<br>5.7.14-google-log
<br>5.1.72-rel14.10

[docs](https://dev.mysql.com/doc)
[github](https://github.com/mysql/mysql-server)

````
Relational DB ⇒ Table ⇒ Row ⇒ Column
````

#### Snippets

````sql
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

````sql
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
````

Horizontal scaling means that you scale by adding more machines into your pool of resources (Replication).
Vertical scaling means that you scale by adding more power (CPU, RAM) to your existing machine.

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
So it is no way to figure out did bot found empty info or didn't find info at all,
unless we use NULL.

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
