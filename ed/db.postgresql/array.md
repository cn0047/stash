Array
-

````sql
CREATE TABLE arrayExample (
    name text,
    num integer array[4],
    tag text[][]
);

INSERT INTO arrayExample VALUES ('one', '{1, 10}', '{"pqsql", "array", "test"}');
INSERT INTO arrayExample VALUES ('two', '{2, 20}', '{{"array", "test"}, {"nested", "multidimensin"}}');
-- Next insert will generate:
-- Multidimensional arrays must have sub-arrays with matching dimensions.
INSERT INTO arrayExample VALUES ('three', array[3, 30], array['pqsql']);
INSERT INTO arrayExample VALUES ('four', array[4], array['1', '2', '3', '4']);
````

````sql
# SELECT:

select * from arrayExample;
 name  |  num   |                  tag
-------+--------+---------------------------------------
 two   | {2,20} | {{array,test},{nested,multidimensin}}
 one   | {1,10} | {pqsql,array,test}
 three | {3,30} | {pqsql}
 four  | {4}    | {1,2,3,4}

select *, tag[1] from arrayExample;
 name  |  num   |                  tag                  |  tag
-------+--------+---------------------------------------+-------
 two   | {2,20} | {{array,test},{nested,multidimensin}} |
 one   | {1,10} | {pqsql,array,test}                    | pqsql
 three | {3,30} | {pqsql}                               | pqsql
 four  | {4}    | {1,2,3,4}                             | 1

select * from arrayExample where tag[1] <>  tag[2];
 name |  num   |        tag
------+--------+--------------------
 one  | {1,10} | {pqsql,array,test}
 four | {4}    | {1,2,3,4}

select * from arrayExample where tag[1][1] <>  tag[2][2];
 name |  num   |                  tag
------+--------+---------------------------------------
 two  | {2,20} | {{array,test},{nested,multidimensin}}

-- slice
select *, tag[2:3] from arrayExample where name = 'four';
 name | num |    tag    |  tag
------+-----+-----------+-------
 four | {4} | {1,2,3,4} | {2,3}

select * from arrayExample where 20 = any (num);
 name |  num   |                   tag
------+--------+------------------------------------------
 two  | {2,20} | {{array,test},{nested,multidimensional}}

-- array has all values equal to 'pqsql'
select * from arrayExample where 'pqsql' = all (tag);
 name  |  num   |   tag
-------+--------+---------
 three | {3,30} | {pqsql}
````

````sql
# UPDATE:

UPDATE arrayExample SET tag[2][2] = 'multidimensional' WHERE name = 'two';
````

````sql
SELECT ARRAY[1, 2, 1+2];
  array
---------
 {1,2,3}

SELECT ARRAY[1,2] || ARRAY[3,4];
 ?column?
-----------
 {1,2,3,4}

SELECT ARRAY[1, 2, 1+2, '4', true]::integer[];
    array
-------------
 {1,2,3,4,1}
````
