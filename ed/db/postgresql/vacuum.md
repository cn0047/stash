VACUUM
-

[docs](https://www.postgresql.org/docs/current/sql-vacuum.html)

VACUUM — garbage-collect and optionally analyze DB.
VACUUM: deletes tuples, prevents transaction ID wraparound, update table statistics.

3 autovacuum workers not enough for modern apps (in 2023).

Recreate index may help to improve vacuum performance.
