create table panel (
  user_id int key,
  email varchar(50),
  name varchar(50)
);

create table panelists (
  user_id int,
  email_active int
);

insert into panel values
(5,              "hello@dummy.com",        "john"),
(6,              "man@city.com",           "Ronn"),
(1,              "fun@dummy.com",          "Sat");

insert into panelists values
(5,              1),
(6,              1),
(1,              1);

select *
from panel t1
left join panelists t2 using(user_id)
where t1.email like "%@dummy.com"
;

+---------+-----------------+------+--------------+
| user_id | email           | name | email_active |
+---------+-----------------+------+--------------+
|       5 | hello@dummy.com | john |            1 |
|       1 | fun@dummy.com   | Sat  |            1 |
+---------+-----------------+------+--------------+

UPDATE panel t1, panelists t2
SET t1.email = "dummy", t2.email_active = 2
WHERE t1.user_id = t2.user_id AND t1.email like "%@dummy.com";
-- OK

-- OR
-- step 1
UPDATE panel SET email = "dummy" WHERE email like "%@dummy.com";

-- step 2
UPDATE panelists t1
JOIN panel t2 on t1.user_id = t2.user_id AND t2.email = "dummy"
SET email_active = 2;
