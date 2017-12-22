Random Stuff
-

````sql
SELECT format('Testing %s, %s, %s, %%', 'one', 'two', 'three');
SELECT format('|%10s|', 'foo');

SELECT ARRAY[1, 2, 1+2];
  array
---------
 {1,2,3}

 SELECT ARRAY[1, 2, 1+2, '4', true]::integer[];
    array
-------------
 {1,2,3,4,1}

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
