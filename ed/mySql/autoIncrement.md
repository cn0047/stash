Auto increment
-

````sql
create table x(id int auto_increment);
-- ERROR 1075 (42000): Incorrect table definition; there can be only one auto column and it must be defined as a key
````
