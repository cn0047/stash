CREATE TABLE parent (
    id INT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;
CREATE TABLE child (
    id INT,
    parent_id INT,
    INDEX par_ind (parent_id),
    FOREIGN KEY (parent_id) REFERENCES parent(id) ON DELETE CASCADE
) ENGINE=INNODB;
INSERT INTO parent VALUES (1), (2), (3);
INSERT INTO child VALUES (1, 1), (2, 2), (3, 3);

SELECT * FROM parent;
+----+
| id |
+----+
|  1 |
|  2 |
|  3 |
+----+
SELECT * FROM child;
+------+-----------+
| id   | parent_id |
+------+-----------+
|    1 |         1 |
|    2 |         2 |
|    3 |         3 |
+------+-----------+

DELETE FROM parent WHERE id = 2;
SELECT * FROM parent;
+----+
| id |
+----+
|  1 |
|  3 |
+----+
SELECT * FROM child;
+------+-----------+
| id   | parent_id |
+------+-----------+
|    1 |         1 |
|    3 |         3 |
+------+-----------+

DROP TABLE parent;
DROP TABLE child;
