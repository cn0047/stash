Random Stuff
-

````sql
SELECT random();

SELECT NOW() - INTERVAL '3 DAY';
SELECT CURRENT_DATE - INTERVAL '4 weeks';

SELECT *
FROM generate_series('2008-03-01 00:00'::timestamp, '2008-03-04 12:00', '10 hours')
;

SELECT format('Testing %s, %s, %s, %%', 'one', 'two', 'three');
SELECT format('|%10s|', 'foo');

SELECT ROW(1,2.5,'this is a test')
;
           row
--------------------------
 (1,2.5,"this is a test")

SELECT * FROM (VALUES (1, 'one'), (2, 'two'), (3, 'three')) AS t
;
 num | letter
-----+--------
   1 | one
   2 | two
   3 | three

SELECT
  count(*) AS unfiltered,
  count(*) FILTER (WHERE i < 5) AS filtered
FROM generate_series(1,10) AS s(i)
;
 unfiltered | filtered
------------+----------
         10 |        4

SELECT EXTRACT(hours from '180 minutes'::interval);
 date_part
-----------
         3

SELECT 'abc' SIMILAR TO '%(b|d)%'; -- true
SELECT 'abc' ~ 'abc'; -- true
SELECT 'abc' ~ '^a'; -- true
SELECT 'abc' ~ '(b|d)'; -- true
SELECT 'abc' ~ '^(b|c)'; -- false

-- select WITH
WITH
  t1 AS (SELECT * FROM (VALUES (1, 'foo'), (2, 'bar'), (3, 'foo')) as t),
  t2 AS (SELECT column2, SUM(column1) FROM t1 GROUP BY column2)
SELECT *
FROM t2
WHERE sum > 0
;
 column2 | sum
---------+-----
 bar     |   2
 foo     |   4

WITH RECURSIVE t(n) AS (
  SELECT * FROM (VALUES (1, 'foo')) as t
  UNION ALL
  SELECT * FROM (VALUES (2, 'bar')) as t
)
SELECT sum(n) FROM t
;
 sum
-----
   3
````

````sql
SELECT * FROM information_schema.columns;

SELECT current_setting('datestyle');
set_config(setting_name, new_value, is_local);
````
