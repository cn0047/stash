create table dates (
  id int auto_increment,
  t timestamp default CURRENT_TIMESTAMP,
  -- d date default CURRENT_DATE   , -- Don't works at version 5.5.41
  -- d date default CURRENT_DATE() , -- Don't works at version 5.5.41
  -- d date default CURDATE        , -- Don't works at version 5.5.41
  -- d date default CURDATE()      , -- Don't works at version 5.5.41
  primary key (id)
);
drop table dates;
create table dates (
  id int auto_increment,
  T1 timestamp default CURRENT_TIMESTAMP,
  t2 timestamp default CURRENT_TIMESTAMP,
  primary key (id)
);
-- Table cannot contains 2 fields with default value CURRENT_TIMESTAMP
