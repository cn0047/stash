CREATE TABLE tableOrders (
    id       INT NOT NULL AUTO_INCREMENT,
    column_2 VARCHAR(256),
    column_3 VARCHAR(256),
    column_4 INT,
    column_5 VARCHAR(256),
    column_6 INT,
    PRIMARY KEY (id)
) ENGINE=InnoDB;


INSERT INTO tableOrders VALUES
    (null, 'foo',  'foo',   9, 'type1', 1),
    (null, 'bar',  'bar',   8, 'type1', 2),
    (null, 'boo',  'boo',   7, 'type1', 3),
    (null, 'bla',  'bla',   6, 'type0', 4),
    (null, 'woo',  'woo',   5, 'type0', 5),
    (null, 'blo',  'blo',   4, 'type0', 6),
    (null, 'blah', 'blah', 35, 'type0', 4),
    (null, 'blah', 'blah', 35, 'type0', 5),
    (null, 'blah', 'blah', 35, 'type1', 2),
    (null, 'blah', 'blah', 35, 'type1', 1),
    (null, 'blah', 'blah', 34, 'type0', 7),
    (null, 'blah', 'blah', 34, 'type1', 9),
    (null, 'blah', 'blah', 34, 'type0', 6),
    (null, 'blah', 'blah', 34, 'type1', 8),
    (null, 'blah', 'fooh', 34, 'type1', 5),
    (null, 'blah', 'fooh', 34, 'type0', 8),
    (null, 'blah', 'fooh', 34, 'type1', 9),
    (null, 'blah', 'fooh', 34, 'type0', 8)
;

SELECT * FROM tableOrders;
+----+----------+----------+----------+----------+----------+
| id | column_2 | column_3 | column_4 | column_5 | column_6 |
+----+----------+----------+----------+----------+----------+
|  1 | foo      | foo      |        9 | type1    |        1 |
|  2 | bar      | bar      |        8 | type1    |        2 |
|  3 | boo      | boo      |        7 | type1    |        3 |
|  4 | bla      | bla      |        6 | type0    |        4 |
|  5 | woo      | woo      |        5 | type0    |        5 |
|  6 | blo      | blo      |        4 | type0    |        6 |
|  7 | blah     | blah     |       35 | type0    |        4 |
|  8 | blah     | blah     |       35 | type0    |        5 |
|  9 | blah     | blah     |       35 | type1    |        2 |
| 10 | blah     | blah     |       35 | type1    |        1 |
| 11 | blah     | blah     |       34 | type0    |        7 |
| 12 | blah     | blah     |       34 | type1    |        9 |
| 13 | blah     | blah     |       34 | type0    |        6 |
| 14 | blah     | blah     |       34 | type1    |        8 |
| 15 | blah     | fooh     |       34 | type1    |        5 |
| 16 | blah     | fooh     |       34 | type0    |        8 |
| 17 | blah     | fooh     |       34 | type1    |        9 |
| 18 | blah     | fooh     |       34 | type0    |        8 |
+----+----------+----------+----------+----------+----------+

SELECT *
FROM  tableOrders
WHERE column_2 = 'blah' AND column_4 = 34
ORDER BY
    CASE WHEN column_3 = 'blah'
        THEN column_5
        ELSE column_6
    END,
    column_4, column_2, id DESC
;
+----+----------+----------+----------+----------+----------+
| id | column_2 | column_3 | column_4 | column_5 | column_6 |
+----+----------+----------+----------+----------+----------+
| 15 | blah     | fooh     |       34 | type1    |        5 |
| 18 | blah     | fooh     |       34 | type0    |        8 |
| 16 | blah     | fooh     |       34 | type0    |        8 |
| 17 | blah     | fooh     |       34 | type1    |        9 |
| 13 | blah     | blah     |       34 | type0    |        6 |
| 11 | blah     | blah     |       34 | type0    |        7 |
| 14 | blah     | blah     |       34 | type1    |        8 |
| 12 | blah     | blah     |       34 | type1    |        9 |
+----+----------+----------+----------+----------+----------+

DROP TABLE tableOrders;
