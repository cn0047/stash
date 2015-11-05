-- MySQL NULLIF() returns NULL when the first is equal to the second expression,
-- other wise it returns the first expression.

SELECT NULLIF(2, 2);
+--------------+
| NULLIF(2, 2) |
+--------------+
|         NULL |
+--------------+
SELECT NULLIF(1, 2);
+--------------+
| NULLIF(1, 2) |
+--------------+
|            1 |
+--------------+

SELECT NULLIF(LEAST("9", "11"), "11");
+--------------------------------+
| NULLIF(LEAST("9", "11"), "11") |
+--------------------------------+
| NULL                           |
+--------------------------------+
SELECT NULLIF(LEAST("9", "11"), "11") + 0;
+------------------------------------+
| NULLIF(LEAST("9", "11"), "11") + 0 |
+------------------------------------+
|                               NULL |
+------------------------------------+
SELECT NULLIF(LEAST("9", "11"), "12") + 0;
+------------------------------------+
| NULLIF(LEAST("9", "11"), "12") + 0 |
+------------------------------------+
|                                  9 | -- Why 9?
+------------------------------------+
SELECT NULLIF(LEAST("9", "11"), "9") + 0;
+-----------------------------------+
| NULLIF(LEAST("9", "11"), "9") + 0 |
+-----------------------------------+
|                                 9 | - Why not "11"? Why not "9"? Why not NULL?
+-----------------------------------+