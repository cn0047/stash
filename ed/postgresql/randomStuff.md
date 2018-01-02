Random Stuff
-

````sql
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
````
