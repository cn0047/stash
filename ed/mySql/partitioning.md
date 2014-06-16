Partitioning
-

*MySQL 5.5*

The user-selected rule by which the division of data is accomplished is known as a partitioning function
This is known as horizontal partitioning—that is, different rows of a table may be assigned to different physical partitions.
MySQL 5.5 does not support vertical partitioning, in which different columns of a table are assigned to different physical partitions.
You cannot use MyISAM for one partition and InnoDB for another.
MySQL partitioning cannot be used with the MERGE, CSV, or FEDERATED storage engines.

Partitioning Types:

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

####RANGE Partitioning
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
1 row in set (0.00 sec)

EXPLAIN PARTITIONS SELECT * FROM orders WHERE date='2008-06-11';
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
| id | select_type | table  | partitions | type   | possible_keys | key  | key_len | ref  | rows | Extra |
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
|  1 | SIMPLE      | orders | p_2008     | system | NULL          | NULL | NULL    | NULL |    1 |       |
+----+-------------+--------+------------+--------+---------------+------+---------+------+------+-------+
1 row in set (0.00 sec)

EXPLAIN PARTITIONS SELECT * FROM orders WHERE date='2014-06-11';
+----+-------------+--------+------------+------+---------------+------+---------+------+------+-------------+
| id | select_type | table  | partitions | type | possible_keys | key  | key_len | ref  | rows | Extra       |
+----+-------------+--------+------------+------+---------------+------+---------+------+------+-------------+
|  1 | SIMPLE      | orders | p_2009     | ALL  | NULL          | NULL | NULL    | NULL |    2 | Using where |
+----+-------------+--------+------------+------+---------------+------+---------+------+------+-------------+
1 row in set (0.00 sec)
````
####LIST
````sql
PARTITION BY LIST(store_id) (
    PARTITION pNorth VALUES IN (3,5,6,9,17),
    PARTITION pEast VALUES IN (1,2,10,11,19,20),
    PARTITION pWest VALUES IN (4,12,13,14,18),
    PARTITION pCentral VALUES IN (7,8,15,16)
);
````

##COLUMNS Partitioning
* All integer types: TINYINT, SMALLINT, MEDIUMINT, INT (INTEGER), and BIGINT. (This is the same as with partitioning by RANGE and LIST.)
<br> Other numeric data types (such as DECIMAL or FLOAT) are not supported as partitioning columns.

* DATE and DATETIME.
<br> Columns using other data types relating to dates or times are not supported as partitioning columns.

* The following string types: CHAR, VARCHAR, BINARY, and VARBINARY.
<br> TEXT and BLOB columns are not supported as partitioning columns.

####RANGE COLUMNS partitioning
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

####LIST COLUMNS partitioning
````sql
PARTITION BY RANGE COLUMNS(renewal) (
    PARTITION pWeek_1 VALUES LESS THAN('2010-02-09'),
    PARTITION pWeek_2 VALUES LESS THAN('2010-02-15'),
    PARTITION pWeek_3 VALUES LESS THAN('2010-02-22'),
    PARTITION pWeek_4 VALUES LESS THAN('2010-03-01')
);
````

####HASH Partitioning
````sql
PARTITION BY HASH( YEAR(hired) )
PARTITIONS 4;
````

####LINEAR HASH Partitioning
Linear hashing, which differs from regular hashing in that linear hashing utilizes a linear powers-of-two algorithm whereas regular hashing employs the modulus of the hashing function's value.
The advantage in partitioning by linear hash is that the adding, dropping, merging, and splitting of partitions is made much faster.
The disadvantage is that data is less likely to be evenly distributed between partitions.
````sql
PARTITION BY LINEAR HASH( YEAR(hired) )
PARTITIONS 4;
````

####KEY Partitioning
````sql
PARTITION BY KEY()
PARTITIONS 2;

PARTITION BY LINEAR KEY (col1)
PARTITIONS 3;
````

####Subpartitioning
````sql
````

####How MySQL Partitioning Handles NULL
Partitioning in MySQL does nothing to disallow NULL. Even though it is permitted to use NULL as the value of an expression that must otherwise yield an integer.
This means that treatment of NULL varies between partitioning of different types, and may produce behavior which you do not expect.

* Handling of NULL with RANGE partitioning.  If you insert a row into a table partitioned by RANGE such that the column value used to determine the partition is NULL, the row is inserted into the lowest partition.

* Handling of NULL with LIST partitioning.  A table that is partitioned by LIST admits NULL values if and only if one of its partitions is defined using that value-list that contains NULL.
````sql
PARTITION BY LIST(c1) (
    PARTITION p0 VALUES IN (0, 3, 6),
    PARTITION p1 VALUES IN (1, 4, 7),
    PARTITION p3 VALUES IN (NULL)
);
````

* Handling of NULL with HASH and KEY partitioning.  NULL is handled somewhat differently for tables partitioned by HASH or KEY. In these cases, any partition expression that yields a NULL value is treated as though its return value were zero.
