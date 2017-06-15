Nested set model
-

[Source](http://mikehillyer.com/articles/managing-hierarchical-data-in-mysql/).

````

+-----------+
|ELECTRONICS|
+-----------+
|
+--------------------------+
|                          |
+-----------+              +--------------------+
|TELEVISIONS|              |PORTABLE ELECTRONICS|
+-----------+              +--------------------+
|                          |
+-------+------+           +--------------+-------------+
|       |      |           |              |             |
+----+  +---+  +------+    +-----------+  +----------+  +------------+
|TUBE|  |LCD|  |PLASMA|    |MP3 PLAYERS|  |CD PLAYERS|  |2 WAY RADIOS|
+----+  +---+  +------+    +-----------+  +----------+  +------------+
                           |
                           +-----+
                           |FLASH|
                           +-----+

````

Hierarchical data has a parent-child relationship
that is not naturally represented in a relational database table.

Queries using nested sets can be expected to be faster
than queries using a stored procedure to traverse an adjacency list.

Nested sets are very slow for inserts because it requires updating left and right domain values
for all records in the table after the insert.

````sql
CREATE TABLE nested_category (
    category_id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(20) NOT NULL,
    lft INT NOT NULL,
    rgt INT NOT NULL
);

INSERT INTO nested_category VALUES
 (1,'ELECTRONICS',1,20)
,(2,'TELEVISIONS',2,9)
,(3,'TUBE',3,4)
,(4,'LCD',5,6)
,(5,'PLASMA',7,8)
,(6,'PORTABLE ELECTRONICS',10,19)
,(7,'MP3 PLAYERS',11,14)
,(8,'FLASH',12,13)
,(9,'CD PLAYERS',15,16)
,(10,'2 WAY RADIOS',17,18)
;

SELECT * FROM nested_category ORDER BY category_id;
+-------------+----------------------+-----+-----+
| category_id | name                 | lft | rgt |
+-------------+----------------------+-----+-----+
|           1 | ELECTRONICS          |   1 |  20 |
|           2 | TELEVISIONS          |   2 |   9 |
|           3 | TUBE                 |   3 |   4 |
|           4 | LCD                  |   5 |   6 |
|           5 | PLASMA               |   7 |   8 |
|           6 | PORTABLE ELECTRONICS |  10 |  19 |
|           7 | MP3 PLAYERS          |  11 |  14 |
|           8 | FLASH                |  12 |  13 |
|           9 | CD PLAYERS           |  15 |  16 |
|          10 | 2 WAY RADIOS         |  17 |  18 |
+-------------+----------------------+-----+-----+
````

#### Retrieving a full tree:

````sql
SELECT node.name
FROM nested_category AS node, nested_category AS parent
WHERE node.lft BETWEEN parent.lft AND parent.rgt AND parent.name = 'ELECTRONICS'
ORDER BY node.lft
;
+----------------------+
| name                 |
+----------------------+
| ELECTRONICS          |
| TELEVISIONS          |
| TUBE                 |
| LCD                  |
| PLASMA               |
| PORTABLE ELECTRONICS |
| MP3 PLAYERS          |
| FLASH                |
| CD PLAYERS           |
| 2 WAY RADIOS         |
+----------------------+
````

#### Finding all the leaf nodes:

````sql
SELECT name FROM nested_category WHERE rgt = lft + 1;
+--------------+
| name         |
+--------------+
| TUBE         |
| LCD          |
| PLASMA       |
| FLASH        |
| CD PLAYERS   |
| 2 WAY RADIOS |
+--------------+
````

#### Retrieving a single path:

````sql
SELECT parent.name
FROM nested_category AS node, nested_category AS parent
WHERE node.lft BETWEEN parent.lft AND parent.rgt AND node.name = 'FLASH'
ORDER BY parent.lft
;
+----------------------+
| name                 |
+----------------------+
| ELECTRONICS          |
| PORTABLE ELECTRONICS |
| MP3 PLAYERS          |
| FLASH                |
+----------------------+
````

#### Finding the depth of the nodes:

````sql
SELECT node.name, (COUNT(parent.name) - 1) AS depth
FROM nested_category AS node, nested_category AS parent
WHERE node.lft BETWEEN parent.lft AND parent.rgt
GROUP BY node.name
ORDER BY node.lft
;
+----------------------+-------+
| name                 | depth |
+----------------------+-------+
| ELECTRONICS          |     0 |
| TELEVISIONS          |     1 |
| TUBE                 |     2 |
| LCD                  |     2 |
| PLASMA               |     2 |
| PORTABLE ELECTRONICS |     1 |
| MP3 PLAYERS          |     2 |
| FLASH                |     3 |
| CD PLAYERS           |     2 |
| 2 WAY RADIOS         |     2 |
+----------------------+-------+

SELECT CONCAT( REPEAT('  ', COUNT(parent.name) - 1), node.name) AS name
FROM nested_category AS node, nested_category AS parent
WHERE node.lft BETWEEN parent.lft AND parent.rgt
GROUP BY node.name
ORDER BY node.lft
;
+------------------------+
| name                   |
+------------------------+
| ELECTRONICS            |
|   TELEVISIONS          |
|     TUBE               |
|     LCD                |
|     PLASMA             |
|   GAME CONSOLES        | -- after add
|   PORTABLE ELECTRONICS |
|     MP3 PLAYERS        |
|       FLASH            |
|     CD PLAYERS         |
|     2 WAY RADIOS       |
+------------------------+
````

#### Adding new nodes:

````sql
LOCK TABLE nested_category WRITE;

SELECT @myRight := rgt FROM nested_category WHERE name = 'TELEVISIONS';

UPDATE nested_category SET rgt = rgt + 2 WHERE rgt > @myRight;
UPDATE nested_category SET lft = lft + 2 WHERE lft > @myRight;

INSERT INTO nested_category(name, lft, rgt) VALUES ('GAME CONSOLES', @myRight + 1, @myRight + 2);

UNLOCK TABLES;
````

#### Deleting nodes:

````sql
LOCK TABLE nested_category WRITE;

SELECT @myLeft := lft, @myRight := rgt, @myWidth := rgt - lft + 1
FROM nested_category
WHERE name = 'GAME CONSOLES';

DELETE FROM nested_category WHERE lft BETWEEN @myLeft AND @myRight;

UPDATE nested_category SET rgt = rgt - @myWidth WHERE rgt > @myRight;
UPDATE nested_category SET lft = lft - @myWidth WHERE lft > @myRight;

UNLOCK TABLES;
````
