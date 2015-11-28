create table author (
  id int,
  name varchar(50),
  primary key (id)
);
create table book (
  id int,
  name varchar(50),
  primary key (id)
);
create table book_autor (
  book_id int,
  author_id int,
  unique key (book_id, author_id),
  foreign key (book_id) references book(id) on delete restrict,
  foreign key (author_id) references author(id) on delete restrict
);

insert into author values (1, 'Ernest'), (2, 'Hemingway'), (3, 'Foo'), (4, 'Boo');
insert into book values
  (1, 'book1'), (2, 'book2'), (3, 'book3'), (4, 'book4'),
  (5, 'book5'), (6, 'book6'), (7, 'book7'), (8, 'book8')
;
insert into book_autor values
  (1, 1), (2, 2), (3, 3), (4, 4),
  (5, 1), (5, 2),
  (6, 1), (6, 2), (6, 3),
  (7, 3), (7, 4)
;

-- select all
select *
from book_autor ba
join author a on ba.author_id = a.id
join book b on ba.book_id = b.id
;
+----+---------+-----------+----+-----------+----+-------+
| id | book_id | author_id | id | name      | id | name  |
+----+---------+-----------+----+-----------+----+-------+
|  1 |       1 |         1 |  1 | Ernest    |  1 | book1 |
|  2 |       1 |         2 |  2 | Hemingway |  1 | book1 |
|  3 |       2 |         3 |  3 | Foo       |  2 | book2 |
|  4 |       3 |         4 |  4 | Boo       |  3 | book3 |
+----+---------+-----------+----+-----------+----+-------+
select *
from book_autor ba
join author a on ba.author_id = a.id and a.name in ('Hemingway')
join book_autor ba2 on ba.book_id = ba2.book_id
join author a2 on ba2.author_id = a2.id and a.id <> a2.id and a2.name in ('Ernest')
;
+----+---------+-----------+----+-----------+----+---------+-----------+----+--------+
| id | book_id | author_id | id | name      | id | book_id | author_id | id | name   |
+----+---------+-----------+----+-----------+----+---------+-----------+----+--------+
|  2 |       1 |         2 |  2 | Hemingway |  1 |       1 |         1 |  1 | Ernest |
+----+---------+-----------+----+-----------+----+---------+-----------+----+--------+

select *
from book_autor ba
join author a on ba.author_id = a.id and a.name = 'Hemingway' and ba.book_id in (
  select book_id
  from book_autor ba
  join author a on ba.author_id = a.id and a.name = 'Ernest'
)
;

-- V2 for Doctrine
create table author (
  id int not null auto_increment key,
  f_name varchar(50) not null default '',
  l_name varchar(50) not null default '',
  dob timestamp not null default '0000-00-00 00:00:00',
  country varchar(50) not null default '',
  unique key (f_name, l_name)
);
create table book (
  id int not null auto_increment key,
  name varchar(50),
  published_at timestamp not null default '0000-00-00 00:00:00',
  -- Unknown database type enum requested, Doctrine\DBAL\Platforms\MySqlPlatform may not support it.
  -- type enum('book', 'serie') not null default 'book',
  type tinyint unsigned not null default 0,
  price decimal not null default '0.00',
  unique key (name)
);
create table book_autor (
-- Table book_autor has no primary key. Doctrine does not support reverse engineering from tables that don't have a primary key.
  id int not null auto_increment key,
  book_id int,
  author_id int,
  unique key (book_id, author_id),
  foreign key (book_id) references book(id) on delete restrict,
  foreign key (author_id) references author(id) on delete restrict
);
