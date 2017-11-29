Performance optimization
-

`SHOW STATUS LIKE 'Last_Query_Cost'`;

#### Improve query performance

`... WHERE duration = 6 OR length = 10;`
Mysql converts `OR` to `UNION` (if not - we have to rewrite query and use union)
and can use 2 indexes (1st for duration, 2nd for length), hence we can create 2 indexes.

`... WHERE duration = 6 AND length = 10;`
It's possible to create 2 separated indexes and mysql will use both.
Or 1 compound index (duration, length).

`... WHERE rightPartOfCompoundIndex ORDER by leftPartOfCompoundIndex` - compound index will be used!

`... WHERE id <> 2` because of `<>` mysql won't use index(
If table have unique id values 1, 2 - we can use `... WHERE id = 1`.

`ORDER BY` one of the most expensive operation for mysql...

`SELECT * ... WHERE length < 100` if result set is big and selects `*` mysql may not use idx_length index
(index created on field length).
It is possible to split query into 2 and use join wiht purpose to reduce result set in each query
hence mysql can use index).

`INNER JOIN table1 INNER JOIN table2 INNER JOIN table3` - order of tables in INNER join section won't affect performance.
Different orders will produce same result and same query cost.

`LEFT JOIN table1 LEFT JOIN table2 LEFT JOIN table3` - order of tables in LEFT join section may affect performance.

`EXISTS` may be faster than `IN` or `JOIN` but it depends...

`ORDER BY ... LIMIT 1` may be faster than `MAX` in case you have index on field used in max function.

`SELECT DISTINCT` faster than `SELECT GROUB BY`.

`GROUB BY title, length` can be replaced with `GROUB BY film_id` because film_id is PK
and contains unique value and combination of title and length also unique
and relationship is 1 - 1.

`UNION ALL` in general faster than `UNION` because mysql don't have to care about duplicates.

#### Index hints

* `FORCE`
* `USE` - just suggestion which key mysql can use.
* `IGNORE`
