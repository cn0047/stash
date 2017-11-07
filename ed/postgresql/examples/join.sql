CREATE TABLE table1 (
  id int NOT NULL PRIMARY KEY ,
  value text NOT NULL
);
CREATE TABLE table2 (
  id int NOT NULL PRIMARY KEY ,
  value text NOT NULL
);

INSERT INTO table1 VALUES
(1, 'one'), (2, 'two'), (3, 'three'), (4, 'four'), (5, 'five');

INSERT INTO table2 VALUES
(3, 'three'), (4, 'four'), (5, 'five'), (6, 'six'), (7, 'seven');


SELECT * FROM table1;
 id | value
----+-------
  1 | one
  2 | two
  3 | three
  4 | four
  5 | five

SELECT * FROM table2;
 id | value
----+-------
  3 | three
  4 | four
  5 | five
  6 | six
  7 | seven

SELECT t1.*, t2.* FROM table1 t1 JOIN table2 t2 on t1.id = t2.id;
SELECT t1.*, t2.* FROM table1 t1 INNER JOIN table2 t2 on t1.id = t2.id;
 id | value | id | value
----+-------+----+-------
  3 | three |  3 | three
  4 | four  |  4 | four
  5 | five  |  5 | five

SELECT t1.*, t2.* FROM table1 t1 LEFT JOIN table2 t2 on t1.id = t2.id;
SELECT t1.*, t2.* FROM table1 t1 LEFT OUTER JOIN table2 t2 on t1.id = t2.id;
 id | value | id | value
----+-------+----+-------
  1 | one   |    |
  2 | two   |    |
  3 | three |  3 | three
  4 | four  |  4 | four
  5 | five  |  5 | five

SELECT t1.*, t2.* FROM table1 t1 RIGHT JOIN table2 t2 on t1.id = t2.id;
SELECT t1.*, t2.* FROM table1 t1 RIGHT OUTER JOIN table2 t2 on t1.id = t2.id;
 id | value | id | value
----+-------+----+-------
  3 | three |  3 | three
  4 | four  |  4 | four
  5 | five  |  5 | five
    |       |  6 | six
    |       |  7 | seven

SELECT t1.*, t2.* FROM table1 t1 LEFT JOIN table2 t2 on t1.id = t2.id WHERE t2.id IS NULL;
 id | value | id | value
----+-------+----+-------
  1 | one   |    |
  2 | two   |    |

SELECT t1.*, t2.* FROM table1 t1 CROSS JOIN table2 t2;
 id | value | id | value
----+-------+----+-------
  1 | one   |  3 | three
  1 | one   |  4 | four
  1 | one   |  5 | five
  1 | one   |  6 | six
  1 | one   |  7 | seven
  2 | two   |  3 | three
  2 | two   |  4 | four
  2 | two   |  5 | five
  2 | two   |  6 | six
  2 | two   |  7 | seven
  3 | three |  3 | three
  3 | three |  4 | four
  3 | three |  5 | five
  3 | three |  6 | six
  3 | three |  7 | seven
  4 | four  |  3 | three
  4 | four  |  4 | four
  4 | four  |  5 | five
  4 | four  |  6 | six
  4 | four  |  7 | seven
  5 | five  |  3 | three
  5 | five  |  4 | four
  5 | five  |  5 | five
  5 | five  |  6 | six
  5 | five  |  7 | seven

SELECT t1.*, t2.* FROM table1 t1 FULL OUTER JOIN table2 t2 on t1.id = t2.id;
 id | value | id | value
----+-------+----+-------
  1 | one   |    |
  2 | two   |    |
  3 | three |  3 | three
  4 | four  |  4 | four
  5 | five  |  5 | five
    |       |  6 | six
    |       |  7 | seven
