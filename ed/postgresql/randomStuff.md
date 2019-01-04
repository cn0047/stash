Random Stuff
-

````sql
SELECT NOW() - INTERVAL '3 DAY';

SELECT format('Testing %s, %s, %s, %%', 'one', 'two', 'three');
SELECT format('|%10s|', 'foo');

SELECT ROW(1,2.5,'this is a test');
           row
--------------------------
 (1,2.5,"this is a test")

SELECT * FROM (VALUES (1, 'one'), (2, 'two'), (3, 'three')) AS t
 num | letter
-----+--------
   1 | one
   2 | two
   3 | three

SELECT
  count(*) AS unfiltered,
  count(*) FILTER (WHERE i < 5) AS filtered
FROM generate_series(1,10) AS s(i);
 unfiltered | filtered
------------+----------
         10 |        4
````

````
SELECT * FROM information_schema.columns;
````
