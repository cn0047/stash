create table partindex (
    friend_name character varying(20) default '',
    age int
);

insert into partindex values ('Jim', null);
insert into partindex values ('Jim', null);

select * from partindex;
 friend_name | age
-------------+-----
 Jim         |
 Jim         |

truncate partindex;

create unique index my_partindex on partindex (friend_name, age) where not (age is null);

insert into partindex values ('Jim', null);
insert into partindex values ('Jim', null);
insert into partindex values ('Jack', 19);
insert into partindex values ('Jack', 19); -- ERROR:  duplicate key
