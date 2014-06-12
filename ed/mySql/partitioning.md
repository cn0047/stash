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


1. RANGE
PARTITION BY RANGE (store_id) (
PARTITION p0 VALUES LESS THAN (10),
PARTITION p1 VALUES LESS THAN (20),
PARTITION p3 VALUES LESS THAN (30)
);

2. LIST
PARTITION BY LIST(store_id) (
PARTITION pNorth VALUES IN (3,5,6,9,17),
PARTITION pEast VALUES IN (1,2,10,11,19,20)
)

3. HASH
PARTITION BY HASH(store_id)
PARTITIONS 4;

4. KEY
PARTITION BY KEY(s1)
PARTITIONS 10;











[>>>](http://dev.mysql.com/doc/refman/5.5/en/partitioning-range.html)