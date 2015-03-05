create table updateTable (
    id int,
    name char(1),
    primary key (id)
) engine=innodb;

drop table updateTable;

insert into updateTable values
(1    ,   'C'),
(2    ,   'A'),
(3    ,   'A'),
(4    ,   'B'),
(5    ,   'B'),
(6    ,   'B'),
(7    ,   'B'),
(8    ,   'X'),
(9    ,   'X'),
(10   ,   'A'),
(11   ,   'A'),
(12   ,   'A'),
(13   ,   'X'),
(14   ,   'X'),
(15   ,   'B'),
(16   ,   'C'),
(17   ,   'C'),
(18   ,   'X'),
(19   ,   'A'),
(20   ,   'A');

select * from updateTable;

/*
Need to receive.

id   |   name
-------------
1    |   C
2    |   A
3    |   NULL
4    |   B
5    |   NULL
6    |   NULL
7    |   NULL
8    |   X
9    |   NULL
10   |   A
11   |   NULL
12   |   NULL
13   |   X
14   |   NULL
15   |   B
16   |   C
17   |   NULL
18   |   X
19   |   A
20   |   NULL
*/

select *
from updateTable t1
left join updateTable t2 on t1.name = t2.name and t1.id+1 = t2.id
;
+----+------+------+------+
| id | name | id   | name |
+----+------+------+------+
|  1 | C    | NULL | NULL |
|  2 | A    |    3 | A    |
|  3 | A    | NULL | NULL |
|  4 | B    |    5 | B    |
|  5 | B    |    6 | B    |
|  6 | B    |    7 | B    |
|  7 | B    | NULL | NULL |
|  8 | X    |    9 | X    |
|  9 | X    | NULL | NULL |
| 10 | A    |   11 | A    |
| 11 | A    |   12 | A    |
| 12 | A    | NULL | NULL |
| 13 | X    |   14 | X    |
| 14 | X    | NULL | NULL |
| 15 | B    | NULL | NULL |
| 16 | C    |   17 | C    |
| 17 | C    | NULL | NULL |
| 18 | X    | NULL | NULL |
| 19 | A    |   20 | A    |
| 20 | A    | NULL | NULL |
+----+------+------+------+

update updateTable set name = null where id in (
    select t2.id
    from updateTable t1
    left join updateTable t2 on t1.name = t2.name and t1.id+1 = t2.id
    where t2.id is not null
);
/*
ERROR 1093 (HY000): You can't specify target table 'updateTable' for update in FROM clause
*/

create temporary table updateTableTmp (
    id int,
    primary key (id)
) engine=innodb;

insert into updateTableTmp
select t2.id
from updateTable t1
left join updateTable t2 on t1.name = t2.name and t1.id+1 = t2.id
where t2.id is not null
;

update updateTable set name = null where id in (
    select id from updateTableTmp
);

select * from updateTable;
+----+------+
| id | name |
+----+------+
|  1 | C    |
|  2 | A    |
|  3 | NULL |
|  4 | B    |
|  5 | NULL |
|  6 | NULL |
|  7 | NULL |
|  8 | X    |
|  9 | NULL |
| 10 | A    |
| 11 | NULL |
| 12 | NULL |
| 13 | X    |
| 14 | NULL |
| 15 | B    |
| 16 | C    |
| 17 | NULL |
| 18 | X    |
| 19 | A    |
| 20 | NULL |
+----+------+
