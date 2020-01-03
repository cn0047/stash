update tbl1 set code = 404 where code in (select code from tbl1);
-- ERROR 1093 (HY000): Table 'tbl1' is specified twice, both as a target for 'UPDATE' and as a separate source for data
