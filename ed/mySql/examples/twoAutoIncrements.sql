create table twoAutoIncrements (
    id1 int auto_increment,
    id2 int auto_increment,
    primary key (id1)
);
ERROR 1075 (42000): Incorrect table definition; there can be only one auto column and it must be defined as a key
