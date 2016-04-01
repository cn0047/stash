
create table one (id int, descr text);
create table two (id int, descr text);
insert into one values
(1, 'o'),
(2, 'tw'),
(3, 'tr'),
(4, 'fr'),
(5, 'fi');

insert into two values
(1, 'ONE'),
(2, 'TWO'),
(5, 'FIVE');

update one join two on one.id = two.id set one.descr = two.descr;

select * from one;
+------+-------+
| id   | descr |
+------+-------+
|    1 | ONE   |
|    2 | TWO   |
|    3 | tr    |
|    4 | fr    |
|    5 | FIVE  |
+------+-------+
