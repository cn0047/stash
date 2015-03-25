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
