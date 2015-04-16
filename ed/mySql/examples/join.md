JOIN
-

````sql
CREATE TABLE `table1` (
  `id` int(11) NOT NULL DEFAULT 0,
  `value` text not null DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

CREATE TABLE `table2` (
  `id` int(11) NOT NULL DEFAULT 0,
  `value` text not null DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;

INSERT INTO table1 VALUES
(1, 'one'), (2, 'two'), (3, 'three'), (4, 'four'), (5, 'five');

INSERT INTO table2 VALUES
(3, 'three'), (4, 'four'), (5, 'five'), (6, 'six'), (7, 'seven');



SELECT * FROM table1;
+----+-------+
| id | value |
+----+-------+
|  1 | one   |
|  2 | two   |
|  3 | three |
|  4 | four  |
|  5 | five  |
+----+-------+
SELECT * FROM table2;
+----+-------+
| id | value |
+----+-------+
|  3 | three |
|  4 | four  |
|  5 | five  |
|  6 | six   |
|  7 | seven |
+----+-------+



SELECT t1.*, t2.* FROM table1 t1 JOIN table2 t2 on t1.id = t2.id;
SELECT t1.*, t2.* FROM table1 t1 INNER JOIN table2 t2 on t1.id = t2.id;
SELECT t1.*, t2.* FROM table1 t1 CROSS JOIN table2 t2 on t1.id = t2.id;
SELECT t1.*, t2.* FROM table1 t1 STRAIGHT_JOIN table2 t2 on t1.id = t2.id;
+----+-------+----+-------+
| id | value | id | value |
+----+-------+----+-------+
|  3 | three |  3 | three |
|  4 | four  |  4 | four  |
|  5 | five  |  5 | five  |
+----+-------+----+-------+




SELECT t1.*, t2.* FROM table1 t1 LEFT JOIN table2 t2 on t1.id = t2.id;
+----+-------+------+-------+
| id | value | id   | value |
+----+-------+------+-------+
|  1 | one   | NULL | NULL  |
|  2 | two   | NULL | NULL  |
|  3 | three |    3 | three |
|  4 | four  |    4 | four  |
|  5 | five  |    5 | five  |
+----+-------+------+-------+
SELECT t1.*, t2.* FROM table1 t1 LEFT OUTER JOIN table2 t2 on t1.id = t2.id;
+----+-------+------+-------+
| id | value | id   | value |
+----+-------+------+-------+
|  1 | one   | NULL | NULL  |
|  2 | two   | NULL | NULL  |
|  3 | three |    3 | three |
|  4 | four  |    4 | four  |
|  5 | five  |    5 | five  |
+----+-------+------+-------+
SELECT t1.*, t2.* FROM table1 t1 RIGHT JOIN table2 t2 on t1.id = t2.id;
+------+-------+----+-------+
| id   | value | id | value |
+------+-------+----+-------+
|    3 | three |  3 | three |
|    4 | four  |  4 | four  |
|    5 | five  |  5 | five  |
| NULL | NULL  |  6 | six   |
| NULL | NULL  |  7 | seven |
+------+-------+----+-------+
SELECT t1.*, t2.* FROM table1 t1 LEFT JOIN table2 t2 on t1.id = t2.id WHERE t2.id IS NULL;
+----+-------+------+-------+
| id | value | id   | value |
+----+-------+------+-------+
|  1 | one   | NULL | NULL  |
|  2 | two   | NULL | NULL  |
+----+-------+------+-------+
````
