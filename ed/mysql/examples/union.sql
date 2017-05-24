SELECT 1 FROM moo LIMIT 1
UNION ALL
SELECT 1 FROM hru LIMIT 1;
+---+
| 1 |
+---+
| 1 |
+---+
1 row in set
-- Because second limit belongs to all query.
