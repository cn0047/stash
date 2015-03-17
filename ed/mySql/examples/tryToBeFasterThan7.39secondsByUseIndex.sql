SELECT count(distinct(users.id)) AS count_1
FROM users USE index (country_2)
WHERE users.country = 'US'
  AND users.last_active > '2015-02-26';
1 row in set (0.35 sec)

SELECT count(distinct(users.id)) AS count_1
FROM offers_clicks USE index (user_id_3), users USE index (country_2)
WHERE users.country = 'US'
  AND users.last_active > '2015-02-26'
  AND offers_clicks.user_id = users.id
  AND offers_clicks.date > '2015-02-14'
  AND offers_clicks.ranking_score < 3.49
  AND offers_clicks.ranking_score > 0.24;
1 row in set (7.39 sec)

explain SELECT count(distinct(users.id)) AS count_1 FROM offers_clicks USE index (user_id_3), users USE index (country_2) WHERE users.country IN ('US') AND users.last_active > '2015-02-26' AND offers_clicks.user_id = users.id AND offers_clicks.date > '2015-02-14' AND offers_clicks.ranking_score < 3.49 AND offers_clicks.ranking_score > 0.24;
+----+-------------+---------------+-------+---------------+-----------+---------+------------------------------+--------+--------------------------+
| id | select_type | table         | type  | possible_keys | key       | key_len | ref                          | rows   | Extra                    |
+----+-------------+---------------+-------+---------------+-----------+---------+------------------------------+--------+--------------------------+
|  1 | SIMPLE      | users         | range | country_2     | country_2 | 14      | NULL                         | 245014 | Using where; Using index |
|  1 | SIMPLE      | offers_clicks | ref   | user_id_3     | user_id_3 | 4       | dejong_pointstoshop.users.id | 270153 | Using where; Using index |
+----+-------------+---------------+-------+---------------+-----------+---------+------------------------------+--------+--------------------------+

users schema:
+---------------------------------+---------------+------+-----+---------------------+----------------+
| Field                           | Type          | Null | Key | Default             | Extra          |
+---------------------------------+---------------+------+-----+---------------------+----------------+
| id                              | int(11)       | NO   | PRI | NULL                | auto_increment |
| country                         | char(2)       | NO   | MUL |                     |                |
| last_active                     | datetime      | NO   | MUL | 2000-01-01 00:00:00 |                |
+---------------+------------+-----------------------------+--------------+-----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| Table         | Non_unique | Key_name                    | Seq_in_index | Column_name     | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
+---------------+------------+-----------------------------+--------------+-----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| users         |          1 | country_2                   |            1 | country         | A         |          14 |     NULL | NULL   |      | BTREE      |         |               |
| users         |          1 | country_2                   |            2 | last_active     | A         |     8048529 |     NULL | NULL   |      | BTREE      |         |               |
+---------------+------------+-----------------------------+--------------+-----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
create table users (
    id int auto_increment,
    country char(2),
    last_active datetime,
    primary key (id),
    key country_2 (country, last_active)
);

clicks schema:
+-----------------+------------------+------+-----+---------------------+----------------+
| Field           | Type             | Null | Key | Default             | Extra          |
+-----------------+------------------+------+-----+---------------------+----------------+
| id              | int(11)          | NO   | PRI | NULL                | auto_increment |
| user_id         | int(11)          | NO   | MUL | 0                   |                |
| offer_id        | int(11) unsigned | NO   | MUL | NULL                |                |
| date            | datetime         | NO   | MUL | 0000-00-00 00:00:00 |                |
| ranking_score   | decimal(5,2)     | NO   | MUL | 0.00                |                |
+---------------+------------+-----------------------------+--------------+-----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| Table         | Non_unique | Key_name                    | Seq_in_index | Column_name     | Collation | Cardinality | Sub_part | Packed | Null | Index_type | Comment | Index_comment |
+---------------+------------+-----------------------------+--------------+-----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
| offers_clicks |          1 | user_id_3                   |            1 | user_id         | A         |         198 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_3                   |            2 | ranking_score   | A         |         198 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_3                   |            3 | date            | A         |         198 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_2                   |            1 | user_id         | A         |    17838712 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_2                   |            2 | date            | A         |    53516137 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_4                   |            1 | user_id         | A         |         198 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_4                   |            2 | date            | A         |         198 |     NULL | NULL   |      | BTREE      |         |               |
| offers_clicks |          1 | user_id_4                   |            3 | ranking_score   | A         |         198 |     NULL | NULL   |      | BTREE      |         |               |
+---------------+------------+-----------------------------+--------------+-----------------+-----------+-------------+----------+--------+------+------------+---------+---------------+
create table offers_clicks (
    id int auto_increment,
    user_id int,
    offer_id int unsigned,
    `date` datetime,
    ranking_score decimal(5,2),
    primary key (id),
    key user_id_3 (user_id, ranking_score, `date`),
    key user_id_2 (user_id, `date`),
    key user_id_4 (user_id, `date`, ranking_score)
);

EXPLAIN
SELECT count(distinct(users.id)) AS count_1
FROM users USE INDEX (country_2)
JOIN offers_clicks USE INDEX (user_id_2)
    ON offers_clicks.user_id = users.id
    AND offers_clicks.date > '2015-02-14'
    AND offers_clicks.ranking_score < 3.49
    AND offers_clicks.ranking_score > 0.24
WHERE users.country = 'US' AND users.last_active > '2015-02-26'
;
