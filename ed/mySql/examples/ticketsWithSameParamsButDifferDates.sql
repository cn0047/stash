create table tickets (
    id int,
    operator_id varchar(10),
    product_id varchar(10),
    created_timestamp timestamp,
    primary key (id)
) engine = innodb;

insert into tickets values
( 1, 'STAFF001', 'acc001', '2015-01-01 22:00:00'),
( 2, 'STAFF003', 'acc004', '2015-11-01 22:00:00'),
( 3, 'STAFF002', 'acc002', '2015-01-01 22:00:00'),
( 4, 'STAFF002', 'acc003', '2015-11-01 22:00:00'),
( 5, 'STAFF001', 'acc005', '2015-01-01 22:00:00'),
( 6, 'STAFF005', 'acc002', '2015-11-01 22:00:00'),
( 7, 'STAFF004', 'acc001', '2015-01-01 22:00:00'),
( 8, 'STAFF001', 'acc001', '2015-12-05 22:00:00'),
( 9, 'STAFF003', 'acc001', '2015-01-01 22:00:00'),
(10, 'STAFF002', 'acc007', '2015-11-01 22:00:00'),
(11, 'STAFF001', 'acc001', '2015-12-03 22:00:00'),
(12, 'STAFF001', 'acc001', '2015-12-01 22:00:00'),
(13, 'STAFF005', 'acc001', '2015-01-01 22:00:00'),
(14, 'STAFF006', 'acc001', '2015-12-01 22:00:00');

select * from tickets;
+----+-------------+------------+---------------------+
| id | operator_id | product_id | created_timestamp   |
+----+-------------+------------+---------------------+
|  1 | STAFF001    | acc001     | 2015-01-01 22:00:00 |
|  2 | STAFF003    | acc004     | 2015-11-01 22:00:00 |
|  3 | STAFF002    | acc002     | 2015-01-01 22:00:00 |
|  4 | STAFF002    | acc003     | 2015-11-01 22:00:00 |
|  5 | STAFF001    | acc005     | 2015-01-01 22:00:00 |
|  6 | STAFF005    | acc002     | 2015-11-01 22:00:00 |
|  7 | STAFF004    | acc001     | 2015-01-01 22:00:00 |
|  8 | STAFF001    | acc001     | 2015-12-05 22:00:00 |
|  9 | STAFF003    | acc001     | 2015-01-01 22:00:00 |
| 10 | STAFF002    | acc007     | 2015-11-01 22:00:00 |
| 11 | STAFF001    | acc001     | 2015-12-03 22:00:00 |
| 12 | STAFF001    | acc001     | 2015-12-01 22:00:00 |
| 13 | STAFF005    | acc001     | 2015-01-01 22:00:00 |
| 14 | STAFF006    | acc001     | 2015-12-01 22:00:00 |
+----+-------------+------------+---------------------+

select *
from tickets t1
join tickets t2 on
    t1.id <> t2.id
    and t1.operator_id = t2.operator_id
    and t1.product_id = t2.product_id
    and ABS(DATEDIFF(t1.created_timestamp, t2.created_timestamp)) <= 2
;
+----+-------------+------------+---------------------+----+-------------+------------+---------------------+
| id | operator_id | product_id | created_timestamp   | id | operator_id | product_id | created_timestamp   |
+----+-------------+------------+---------------------+----+-------------+------------+---------------------+
| 11 | STAFF001    | acc001     | 2015-12-03 22:00:00 |  8 | STAFF001    | acc001     | 2015-12-05 22:00:00 |
|  8 | STAFF001    | acc001     | 2015-12-05 22:00:00 | 11 | STAFF001    | acc001     | 2015-12-03 22:00:00 |
| 12 | STAFF001    | acc001     | 2015-12-01 22:00:00 | 11 | STAFF001    | acc001     | 2015-12-03 22:00:00 |
| 11 | STAFF001    | acc001     | 2015-12-03 22:00:00 | 12 | STAFF001    | acc001     | 2015-12-01 22:00:00 |
+----+-------------+------------+---------------------+----+-------------+------------+---------------------+
