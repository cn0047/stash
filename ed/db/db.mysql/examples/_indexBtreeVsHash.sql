CREATE TABLE `lookup` (
  `id` int(11) NOT NULL DEFAULT '0',
  `location1` varchar(50) DEFAULT NULL,
  `location2` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `id_index_1` (`location1`) USING BTREE,
  KEY `id_index_2` (`location2`) USING BTREE
) ENGINE=InnoDB;

SHOW INDEX FROM lookup;
+--------+------------+------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| Table  | Non_unique | Key_name   | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
+--------+------------+------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| lookup |          0 | PRIMARY    |            1 | id          | A         |           0 |     NULL | NULL   |      | BTREE      |         |               |
| lookup |          1 | id_index_1 |            1 | location1   | A         |           0 |     NULL | NULL   | YES  | BTREE      |         |               |
| lookup |          1 | id_index_2 |            1 | location2   | A         |           0 |     NULL | NULL   | YES  | BTREE      |         |               |
+--------+------------+------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+

DROP TABLE lookup;

CREATE TABLE lookup (
  id int(11) NOT NULL DEFAULT '0',
  location1 varchar(50) DEFAULT NULL,
  location2 varchar(50) DEFAULT NULL,
  PRIMARY KEY (id),
  KEY id_index_1 (location1) USING BTREE,
  KEY id_index_2 (location2) USING HASH
) ENGINE=MEMORY;

SHOW INDEX FROM lookup;
+--------+------------+------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| Table  | Non_unique | Key_name   | Seq_in_index | Column_name | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
+--------+------------+------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| lookup |          0 | PRIMARY    |            1 | id          | NULL      |           0 |     NULL | NULL   |      | HASH       |         |               |
| lookup |          1 | id_index_1 |            1 | location1   | A         |        NULL |     NULL | NULL   | YES  | BTREE      |         |               |
| lookup |          1 | id_index_2 |            1 | location2   | NULL      |           0 |     NULL | NULL   | YES  | HASH       |         |               |
+--------+------------+------------+--------------+-------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
