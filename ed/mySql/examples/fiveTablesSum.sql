create table fiveTablesSum_A (
    id_A char(2),
    name varchar(10),
    primary key (id_A)
)engine = innodb;
insert into fiveTablesSum_A values
('A1', 'Name A'),
('A2', 'Name b');

create table fiveTablesSum_B (
    id int auto_increment,
    id_B char(2),
    id_A char(2),
    primary key (id)
)engine = innodb;
insert into fiveTablesSum_B values
(null, 'B1', 'A1'),
(null, 'B2', 'A2');

create table fiveTablesSum_C (
    id_C char(2),
    id_B char(2),
    valueC int,
    primary key (id_C)
)engine = innodb;
insert into fiveTablesSum_C values
('C1', 'B1', 1),
('C2', 'B2', 1),
('C3', 'B1', 1);

create table fiveTablesSum_D (
    id_D char(2),
    id_A char(2),
    valueD int,
    primary key (id_D)
)engine = innodb;
insert into fiveTablesSum_D values
('D1', 'A1', 1000),
('D2', 'A1', 500),
('D3', 'A1', 1000),
('D4', 'A2', 1000),
('D5', 'A2', 500);

create table fiveTablesSum_E (
    id_E char(2),
    id_A char(2),
    valueE int,
    primary key (id_E)
)engine = innodb;
insert into fiveTablesSum_E values
('E1', 'A1', 2000),
('E2', 'A1', 1500),
('E3', 'A2', 500),
('E4', 'A2', 500),
('E5', 'A2', 1000);

drop table fiveTablesSum_A;
drop table fiveTablesSum_B;
drop table fiveTablesSum_C;
drop table fiveTablesSum_D;
drop table fiveTablesSum_E;

select * from fiveTablesSum_A;
select * from fiveTablesSum_B;
select * from fiveTablesSum_C;
select * from fiveTablesSum_D;
select * from fiveTablesSum_E;




select tb.id_A id_A, sum(valueC) tot_C
from fiveTablesSum_C tc
join fiveTablesSum_B tb on tc.id_B = tb.id_B
group by id_A;
+------+-------+
| id_A | tot_C |
+------+-------+
| A1   |     2 |
| A2   |     1 |
+------+-------+

select id_A id_A, sum(valueD) tot_D from fiveTablesSum_D group by id_A;
+------+-------+
| id_A | tot_D |
+------+-------+
| A1   |  2500 |
| A2   |  1500 |
+------+-------+

select id_A id_A, sum(valueE) tot_E from fiveTablesSum_E group by id_A;
+------+-------+
| id_A | tot_E |
+------+-------+
| A1   |  3500 |
| A2   |  2000 |
+------+-------+

select
    t1.id_A,
    t2.tot_C, t3.tot_D, t4.tot_E,
    t2.tot_C + t3.tot_D + t4.tot_E tot_ALL
from fiveTablesSum_A t1
left join (
    select tb.id_A id_A, sum(valueC) tot_C
    from fiveTablesSum_C tc
    join fiveTablesSum_B tb on tc.id_B = tb.id_B
    group by id_A
) t2 on t1.id_A = t2.id_A
left join (
    select id_A id_A, sum(valueD) tot_D from fiveTablesSum_D group by id_A
) t3 on t1.id_A = t3.id_A
left join (
    select id_A id_A, sum(valueE) tot_E from fiveTablesSum_E group by id_A
) t4 on t1.id_A = t4.id_A
;
+------+-------+-------+-------+---------+
| id_A | tot_C | tot_D | tot_E | tot_ALL |
+------+-------+-------+-------+---------+
| A1   |     2 |  2500 |  3500 |    6002 |
| A2   |     1 |  1500 |  2000 |    3501 |
+------+-------+-------+-------+---------+
