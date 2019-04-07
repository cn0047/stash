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

drop table if exists types;
create table types (
  id serial not null primary key,
  b boolean not null default true,
  c char not null default '',
  c2 char(9) not null default '' unique,
  vc char(9) not null default '',
  si smallint not null default 0,
  i smallint not null default -1,
  s serial not null, -- default value generates ERROR:  multiple default values specified for column "s" of table "types"
  f float not null default 0.0 check (f > -1),
  r real not null default 0.0,
  n numeric not null default 0.0,
  d date not null default now(),
  t time not null default now(),
  ts timestamp not null default now(),
  tz timestamptz not null default now(),
  it interval not null default '1 day 3 hours 31 min 24 sec'
  -- p point,
  -- l line,
  -- l lseg,
  -- b box,
  -- p polygon,
  -- i inet,
  -- m macaddr,
);
insert into types (id) values (1);
select * from types;

-- error: Failing row contains
insert into cttest values (1, -1, 1, 1.1, 'one', 'one', '{}', default);
insert into cttest values (1, 1, 1, 1.1, 'one', 'one', '{}', default);
-- ok
insert into cttest values (1, 1, 2, 1.1, 'one', 'one', '{}', default);
insert into cttest values (2, 2, 3, 1.1, 'one', 'one', default, now());
--
select * from cttest;
