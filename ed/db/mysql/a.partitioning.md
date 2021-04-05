Partitioning
-

<br>MySQL 5.7
<br>MySQL 5.5

The user-selected rule by which the division of data is accomplished is known as a partitioning function.
<br>MySQL 5.5 does not support vertical partitioning, in which different columns of a table are assigned to different physical partitions.
<br>This is known as horizontal partitioning—that is, different rows of a table may be assigned to different physical partitions.

Horizontal sharding - split 1 big table into 2 small by logic rule (leads_0, leads_1, leads_2).
Vertical sharding - move table into separate database.

Subpartitioning — composite partitioning, is further division of each partition in a partitioned table.

## Partitioning Types:

    1 RANGE Partitioning
    2 LIST Partitioning
    3 COLUMNS Partitioning
    4 HASH Partitioning
    5 KEY Partitioning


````sql
ALTER TABLE employees DROP PARTITION p0;
ALTER TABLE employees TRUNCATE PARTITION pWest;
ALTER TABLE employees PARTITION BY RANGE COLUMNS (hired) (
    PARTITION p3 VALUES LESS THAN ('2000-01-01'),
    PARTITION p4 VALUES LESS THAN ('2010-01-01'),
    PARTITION p5 VALUES LESS THAN (MAXVALUE)
);

SELECT PARTITION_NAME,TABLE_ROWS FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_NAME = 'table';

SELECT TABLE_NAME, PARTITION_NAME, TABLE_ROWS, AVG_ROW_LENGTH, DATA_LENGTH
FROM INFORMATION_SCHEMA.PARTITIONS
WHERE TABLE_NAME ='table';
````

#### RANGE Partitioning
````sql
CREATE TABLE orders (
    date DATE,
    note VARCHAR(500)
) ENGINE = MYISAM
PARTITION BY RANGE( YEAR(date) ) (
PARTITION p_old VALUES LESS THAN(2008),
PARTITION p_2008 VALUES LESS THAN(2009),
PARTITION p_2009 VALUES LESS THAN(MAXVALUE)
);

INSERT INTO orders SET date='2014-06-11', note='1';
INSERT INTO orders SET date='2014-06-11', note='2';
INSERT INTO orders SET date='2007-06-11', note='3';
INSERT INTO orders SET date='2008-06-11', note='4';

SELECT * FROM orders;
+------------+------+
| date       | note |
+------------+------+
| 2007-06-11 | 3    |
| 2008-06-11 | 4    |
| 2014-06-11 | 1    |
| 2014-06-11 | 2    |
+------------+------+
4 rows in set (0.00 sec)

EXPLAIN PARTITIONS SELECT * FROM orders WHERE date='2007-06-11';
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
| id | select_type | table  | partitions | type   | possible_keys | key  | key_len | ref  | rows | Extra |
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
|  1 | SIMPLE      | orders | p_old      | system | NULL          | NULL | NULL    | NULL |    1 |       |
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+

EXPLAIN PARTITIONS SELECT * FROM orders WHERE date='2008-06-11';
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
| id | select_type | table  | partitions | type   | possible_keys | key  | key_len | ref  | rows | Extra |
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
|  1 | SIMPLE      | orders | p_2008     | system | NULL          | NULL | NULL    | NULL |    1 |       |
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+

EXPLAIN PARTITIONS SELECT * FROM orders WHERE date='2014-06-11';
+----+-------------+--------+------------+------+---------------+------+---------+------+------+-------------+
| id | select_type | table  | partitions | type | possible_keys | key  | key_len | ref  | rows | Extra       |
+----+-------------+--------+------------+------+---------------+------+---------+------+------+-------------+
|  1 | SIMPLE      | orders | p_2009     | ALL  | NULL          | NULL | NULL    | NULL |    2 | Using where |
+----+-------------+--------+------------+------+---------------+------+---------+------+------+-------------+

CREATE TRIGGER ins_orders BEFORE INSERT ON orders
FOR EACH ROW SET New.note = concat(NEW.note, ' triggered');

INSERT INTO orders SET date='2016-09-01', note='5';

SELECT * FROM orders;
+------------+-------------+
| date       | note        |
+------------+-------------+
| 2007-06-11 | 3           |
| 2008-06-11 | 4           |
| 2014-06-11 | 1           |
| 2014-06-11 | 2           |
| 2016-09-01 | 5 triggered |
+------------+-------------+
````
#### LIST
````sql
PARTITION BY LIST(store_id) (
    PARTITION pNorth VALUES IN (3,5,6,9,17),
    PARTITION pEast VALUES IN (1,2,10,11,19,20),
    PARTITION pWest VALUES IN (4,12,13,14,18),
    PARTITION pCentral VALUES IN (7,8,15,16)
);
````

In case value not present in any list - `ERROR 1526 (HY000): Table has no partition for value 77`.

#### COLUMNS Partitioning
* All integer types: TINYINT, SMALLINT, MEDIUMINT, INT (INTEGER), and BIGINT. (This is the same as with partitioning by RANGE and LIST.)
<br> Other numeric data types (such as DECIMAL or FLOAT) are not supported as partitioning columns.

* DATE and DATETIME.
<br> Columns using other data types relating to dates or times are not supported as partitioning columns.

* The following string types: CHAR, VARCHAR, BINARY, and VARBINARY.
<br> TEXT and BLOB columns are not supported as partitioning columns.

#### RANGE COLUMNS partitioning
````sql
PARTITION BY RANGE COLUMNS(a, b) (
    PARTITION p0 VALUES LESS THAN (5, 12),
    PARTITION p3 VALUES LESS THAN (MAXVALUE, MAXVALUE)
);

SELECT (5,10) < (5,12), (5,11) < (5,12), (5,12) < (5,12);
+-----------------+-----------------+-----------------+
| (5,10) < (5,12) | (5,11) < (5,12) | (5,12) < (5,12) |
+-----------------+-----------------+-----------------+
|               1 |               1 |               0 |
+-----------------+-----------------+-----------------+
SELECT (0,25,50) < (10,20,100), (10,20,100) < (10,30,50), (0,25,50) < (20,20,100), (20,20,100) < (10,30,50);
+-------------------------+--------------------------+-------------------------+--------------------------+
| (0,25,50) < (10,20,100) | (10,20,100) < (10,30,50) | (0,25,50) < (20,20,100) | (20,20,100) < (10,30,50) |
+-------------------------+--------------------------+-------------------------+--------------------------+
|                       1 |                        1 |                       1 |                        0 |
+-------------------------+--------------------------+-------------------------+--------------------------+
````

#### LIST COLUMNS partitioning
````sql
PARTITION BY RANGE COLUMNS(renewal) (
    PARTITION pWeek_1 VALUES LESS THAN('2010-02-09'),
    PARTITION pWeek_2 VALUES LESS THAN('2010-02-15'),
    PARTITION pWeek_3 VALUES LESS THAN('2010-02-22'),
    PARTITION pWeek_4 VALUES LESS THAN('2010-03-01')
);
````

#### HASH Partitioning
````sql
PARTITION BY HASH( YEAR(hired) )
PARTITIONS 4;
````

#### LINEAR HASH Partitioning
Linear hashing, which differs from regular hashing
in that linear hashing utilizes a linear powers-of-two algorithm
whereas regular hashing employs the modulus of the hashing function&#39;s value.
The advantage in partitioning by linear hash
is that the adding, dropping, merging, and splitting of partitions is made much faster.
The disadvantage is that data is less likely to be evenly distributed between partitions.
````sql
PARTITION BY LINEAR HASH( YEAR(hired) )
PARTITIONS 4;
````

#### KEY Partitioning
````sql
PARTITION BY KEY()
PARTITIONS 2;

PARTITION BY LINEAR KEY (col1)
PARTITIONS 3;
````

## How MySQL Partitioning Handles NULL

Partitioning in MySQL does nothing to disallow NULL.
Even though it is permitted to use NULL as the value of an expression that must otherwise yield an integer.
This means that treatment of NULL varies between partitioning of different types,
and may produce behavior which you do not expect.

* Handling of NULL with RANGE partitioning.
  If you insert a row into a table partitioned by RANGE such that the column value used to determine the partition is NULL,
  the row is inserted into the lowest partition.

* Handling of NULL with LIST partitioning.
  A table that is partitioned by LIST admits NULL values
  if and only if one of its partitions is defined using that value-list that contains NULL.

````sql
PARTITION BY LIST(c1) (
    PARTITION p0 VALUES IN (0, 3, 6),
    PARTITION p1 VALUES IN (1, 4, 7),
    PARTITION p3 VALUES IN (NULL)
);
````

* Handling of NULL with HASH and KEY partitioning.
  NULL is handled somewhat differently for tables partitioned by HASH or KEY.
  In these cases, any partition expression that yields a NULL value is treated as though its return value were zero.

## Partition Management

#### Management of RANGE and LIST Partitions
It is very important to remember that, when you drop a partition, you also delete all the data that was stored in that partition.
<br>With tables that are partitioned by range, you can use ADD PARTITION to add new partitions to the high end of the partitions list only.
````sql
ALTER TABLE members ADD PARTITION (PARTITION p3 VALUES LESS THAN (2000));
ALTER TABLE members
    REORGANIZE PARTITION p0 INTO (
        PARTITION n0 VALUES LESS THAN (1960),
        PARTITION n1 VALUES LESS THAN (1970)
);
ALTER TABLE members REORGANIZE PARTITION s0,s1 INTO (
    PARTITION p0 VALUES LESS THAN (1970)
);
````
````sql
ALTER TABLE tt ADD PARTITION (PARTITION p2 VALUES IN (7, 14, 21));
ALTER TABLE tt REORGANIZE PARTITION p1,np INTO (
    PARTITION p1 VALUES IN (6, 18),
    PARTITION np VALUES in (4, 8, 12)
);
````

#### #Management of HASH and KEY Partitions
````sql
CREATE TABLE clients (
    id INT,
    fname VARCHAR(30),
    lname VARCHAR(30),
    signed DATE
)
PARTITION BY HASH( MONTH(signed) )
PARTITIONS 12;

SELECT PARTITION_NAME,TABLE_ROWS FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_NAME = 'clients';
+----------------+------------+
| PARTITION_NAME | TABLE_ROWS |
+----------------+------------+
| p0             |          0 |
| p1             |          0 |
| p2             |          0 |
| p3             |          0 |
| p4             |          0 |
| p5             |          0 |
| p6             |          0 |
| p7             |          0 |
| p8             |          0 |
| p9             |          0 |
| p10            |          0 |
| p11            |          0 |
+----------------+------------+
12 rows in set (0.00 sec)

-- To reduce the number of partitions from twelve to eight
ALTER TABLE clients COALESCE PARTITION 4;

SELECT PARTITION_NAME,TABLE_ROWS FROM INFORMATION_SCHEMA.PARTITIONS WHERE TABLE_NAME = 'clients';
+----------------+------------+
| PARTITION_NAME | TABLE_ROWS |
+----------------+------------+
| p0             |          0 |
| p1             |          0 |
| p2             |          0 |
| p3             |          0 |
| p4             |          0 |
| p5             |          0 |
| p6             |          0 |
| p7             |          0 |
+----------------+------------+
8 rows in set (0.00 sec)

-- To increase the number of partitions for the clients table from 12 to 18
ALTER TABLE clients ADD PARTITION PARTITIONS 6;
````

#### Maintenance of Partitions
Rebuilds the partition; this has the same effect as dropping all records stored in the partition, then reinserting them.
````sql
ALTER TABLE t1 REBUILD PARTITION p0, p1;
````

#### Obtaining Information About Partitions
````sql
SHOW TABLE STATUS FROM dbName where Name = 'tableName'; -- look to Create_options
EXPLAIN PARTITIONS SELECT * FROM clients;
````

## Partition Pruning

The core concept behind partition pruning is
“Do not scan partitions where there can be no matching values”.

MySQL can apply partition pruning to SELECT, DELETE, and UPDATE statements.
INSERT statements currently cannot be pruned.

Pruning can also be applied for tables partitioned on a DATE or DATETIME column
when the partitioning expression uses the YEAR() or TO_DAYS() function.

This optimization is used only if the range size is smaller than the number of partitions.

## Partition Selection

SQL statements supporting explicit partition selection are listed here:

* SELECT
* DELETE
* INSERT
* REPLACE
* UPDATE
* LOAD DATA
* LOAD XML

````sql
SELECT * FROM orders PARTITION(p_old, p_2008) ORDER BY date;
+------------+------+
| date       | note |
+------------+------+
| 2007-06-11 | 3    |
| 2008-06-11 | 4    |
+------------+------+
````

## Restrictions and Limitations on Partitioning

You cannot use MyISAM for one partition and InnoDB for another.

MySQL partitioning cannot be used with the MERGE, CSV, or FEDERATED storage engines.

Temporary tables cannot be partitioned.

Use of **INSERT DELAYED** to insert rows into a partitioned table **is not supported**.

All **columns used in the partitioning expression** for a partitioned table **must be part of every unique key** that the table may have.
This also includes the table's primary key, since it is by definition a unique key.

For a partitioned MyISAM table, MySQL uses 2 file descriptors for each partition, for each such table that is open.

The maximum possible number of partitions for a given table not using the NDB storage engine is 8192.

The **query cache is not supported for partitioned tables**,
and is automatically disabled for queries involving partitioned tables.
The query cache cannot be enabled for such queries.

Foreign keys not supported for partitioned InnoDB tables.
In addition, `ALTER TABLE ... OPTIMIZE PARTITION` does not work correctly
with partitioned tables that use the InnoDB storage engine.
Use `ALTER TABLE ... REBUILD PARTITION`
and `ALTER TABLE ... ANALYZE PARTITION`, instead, for such tables.

Partitioned tables **do not support FULLTEXT indexes**.

Columns with spatial data types such as POINT or GEOMETRY cannot be used in partitioned tables.

A partitioning key must be either an integer column or an expression that resolves to an integer.
Expressions employing **ENUM columns cannot be used**.

Only the MySQL functions shown here are allowed in partitioning expressions:

* ABS, CEILING, FLOOR, MOD
* DATEDIFF, DAY, DAYOFMONTH, DAYOFWEEK, DAYOFYEAR, EXTRACT, HOUR, MICROSECOND, MINUTE, MONTH, QUARTER, SECOND, WEEKDAY, YEAR, YEARWEEK
* TIME_TO_SEC, TO_DAYS, TO_SECONDS, UNIX_TIMESTAMP

A SELECT from a partitioned MyISAM table locks only those partitions actually containing rows
that satisfy the SELECT statement's WHERE condition are locked.
