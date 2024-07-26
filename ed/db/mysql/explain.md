EXPLAIN
-

Example:

````sql
+----+-------------+-------+------------+-------+---------------+---------+---------+------+------+----------+-------------+
| id | select_type | table | partitions | type  | possible_keys | key     | key_len | ref  | rows | filtered | Extra       |
+----+-------------+-------+------------+-------+---------------+---------+---------+------+------+----------+-------------+
|  1 | SIMPLE      | test  | NULL       | index | NULL          | PRIMARY | 4       | NULL |    1 |   100.00 | Using index |
+----+-------------+-------+------------+-------+---------------+---------+---------+------+------+----------+-------------+
````

`key_len` - length of the key that MySQL decided to use.
Determine how many parts of a multiple-part key MySQL actually uses.

`filtered` - estimated percentage of table rows that will be filtered by the table condition.
100.00 - good.

Column type - explains how exactly mysql will perform search.

`ALL` - table scan (worst).
`index` - scan using index (better), if extra has "Using index" - covering index used.
`range` - same as index but with `BETWEEN, >, IN` clauses (good).
`ref` - index lookup, will return rows corresponding to single reference value (very good).
`eq_ref` - for primary or unique keys, case when will be returned single result (fast).
`const, system` - part of query converted to constant, as optimisation (faster).
`NULL` - query solved on optimization phase, at worst index will be used but not table (fastest).
