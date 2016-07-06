select *
from (select 'ok' v, 1 k) t
where t.k < any (select 1 union all select 2)
;
+----+---+
| v  | k |
+----+---+
| ok | 1 |
+----+---+
