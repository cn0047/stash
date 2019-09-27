Relationships between Models
-

[Source](https://docs.phalconphp.com/en/latest/reference/models.html#relationships-between-models).

````sql
insert into robots values
(null, 'robot 1', 'type 1', 2016),
(null, 'robot 2', 'type 2', 2016);

insert into parts values
(null, 'part 1'),
(null, 'part 2'),
(null, 'part 3');

insert into robots_parts values
(null, 1, 1, now()),
(null, 1, 2, now()),
(null, 2, 2, now()),
(null, 2, 3, now());
````
