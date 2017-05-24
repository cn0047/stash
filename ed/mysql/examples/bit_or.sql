create temporary table testBinOperation
select conv(2, 10, 2) as mask
union select conv(5, 10, 2) as mask
union select conv(16, 10, 2) as mask
;
select lpad(mask, 10, 0) from testBinOperation
;
/*
+-------------------+
| lpad(mask, 10, 0) |
+-------------------+
| 0000000010        |
| 0000000101        |
| 0000010000        |
+-------------------+
*/
select bit_or(mask) from testBinOperation
;
/*
+--------------+
| bit_or(mask) |
+--------------+
|        10111 |
+--------------+
*/
