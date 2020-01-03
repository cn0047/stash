CREATE TABLE tableAVGWithNull (id INT DEFAULT NULL) ENGINE=InnoDB;

INSERT INTO tableAVGWithNull VALUES (0), (1), (2), (NULL);

select * from tableAVGWithNull;
+------+
| id   |
+------+
|    0 |
|    1 |
|    2 |
| NULL |
+------+

SELECT AVG(id) FROM tableAVGWithNull;
+---------+
| AVG(id) |
+---------+
|  1.0000 |
+---------+
