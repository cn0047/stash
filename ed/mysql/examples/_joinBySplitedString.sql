create table prod (
    codes int,
    `desc` text,
    primary key (codes)
);
create table prod_details (
    codes int,
    mat_id text,
    primary key (codes)
);
create table materials (
    mat_id int,
    mat_name text,
    primary key (mat_id)
);

insert into prod values
(100  ,'table 1'),
(101  ,'chair 1'),
(102  ,'chair 2');
insert into prod_details values
(100  ,'50,52'),
(101  ,'53'),
(102  ,'51,52,54');
insert into materials values
(50    ,'pine wood'),
(51    ,'acacia wood'),
(52    ,'MDF'),
(53    ,'stainless s'),
(54    ,'leather');

select * from prod;
select * from prod_details;
select * from materials;

+-------+---------+
| codes | desc    |
+-------+---------+
|   100 | table 1 |
|   101 | chair 1 |
|   102 | chair 2 |
+-------+---------+
+-------+----------+
| codes | mat_id   |
+-------+----------+
|   100 | 50,52    |
|   101 | 53       |
|   102 | 51,52,54 |
+-------+----------+
+--------+-------------+
| mat_id | mat_name    |
+--------+-------------+
|     50 | pine wood   |
|     51 | acacia wood |
|     52 | MDF         |
|     53 | stainless s |
|     54 | leather     |
+--------+-------------+

select
    d.*
    ,m.*
from prod_details d
left join materials m on find_in_set(m.mat_id, d.mat_id) > 0
;
+-------+----------+--------+-------------+
| codes | mat_id   | mat_id | mat_name    |
+-------+----------+--------+-------------+
|   100 | 50,52    |     50 | pine wood   |
|   100 | 50,52    |     52 | MDF         |
|   101 | 53       |     53 | stainless s |
|   102 | 51,52,54 |     51 | acacia wood |
|   102 | 51,52,54 |     52 | MDF         |
|   102 | 51,52,54 |     54 | leather     |
+-------+----------+--------+-------------+
