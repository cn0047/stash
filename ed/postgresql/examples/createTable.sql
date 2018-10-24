drop table if exists cttest;

create table cttest (
  id serial not null primary key,
  n int,
  m numeric,
  d double precision,
  -- b boolean,
  s character varying(20) default '',
  t text,
  j json not null default '{}',
  ts timestamp default current_timestamp,
  -- PRIMARY KEY (id, n)
  unique (n, m),
  -- constraints
  check (n > 0),
  check (m > n)
);

-- error: Failing row contains
insert into cttest values (1, -1, 1, 1.1, 'one', 'one', '{}', default);
insert into cttest values (1, 1, 1, 1.1, 'one', 'one', '{}', default);
-- ok
insert into cttest values (1, 1, 2, 1.1, 'one', 'one', '{}', default);
insert into cttest values (2, 2, 3, 1.1, 'one', 'one', default, now());
--
select * from cttest;
