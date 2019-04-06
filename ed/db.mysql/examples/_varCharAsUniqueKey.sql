create table uk(
    id int Auto_increment,
    browser varchar(20),
    primary key (id),
    unique key (browser)
);
insert into uk values (null, 'firefox');
insert into uk values (null, 'Chrome');
insert into uk values (null, 'Firefox');
-- ERROR 1062 (23000): Duplicate entry 'Firefox' for key 'browser'
select * from uk;
+----+---------+
| id | browser |
+----+---------+
|  2 | Chrome  |
|  1 | firefox |
+----+---------+
2 rows in set (0.00 sec)

create table uk2 (
    id int Auto_increment,
    browser varchar(20) CHARACTER SET utf8 COLLATE utf8_bin,
    primary key (id),
    unique key (browser)
);
-- OR
create table uk2 (
    id int Auto_increment,
    browser varchar(20) CHARACTER SET utf8 COLLATE utf8_bin not null unique,
    primary key (id)
);
insert into uk2 values (null, 'firefox');
insert into uk2 values (null, 'Chrome');
insert into uk2 values (null, 'Firefox');
select * from uk2;
+----+---------+
| id | browser |
+----+---------+
|  2 | Chrome  |
|  3 | Firefox |
|  1 | firefox |
+----+---------+
